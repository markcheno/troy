package troy

import (
	"sync"
)

type Write struct {
	Queue []string
	Store Store
}

func (w *Write) V(value string) *Write {
	w.Queue = append(w.Queue, value)
	return w
}

func (w *Write) Out(predicate string) *Write {
	w.Queue = append(w.Queue, predicate)
	return w
}

func (w *Write) Exec() {
	var wg sync.WaitGroup
	for i := 0; i < len(w.Queue); i += 2 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if n+2 >= len(w.Queue) {
				return
			}
			subject := w.Queue[n]
			predicate := w.Queue[n+1]
			object := w.Queue[n+2]
			w.Store.Update(subject, predicate, object)
		}(i)
	}
	wg.Wait()
}
