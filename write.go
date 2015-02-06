package troy

import (
	"sync"
)

type Write struct {
	queue []string
	Store Store
}

func (w *Write) V(value string) *Write {
	w.queue = append(w.queue, value)
	return w
}

func (w *Write) Out(predicate string) *Write {
	w.queue = append(w.queue, predicate)
	return w
}

func (w *Write) Exec() {
	var wg sync.WaitGroup
	for i := 0; i < len(w.queue); i += 2 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if n+3 > len(w.queue) {
				return
			}
			subject := w.queue[n]
			predicate := w.queue[n+1]
			object := w.queue[n+2]
			w.Store.Update(subject, predicate, object)
		}(i)
	}
	wg.Wait()
}
