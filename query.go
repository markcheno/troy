package troy

import (
	"sync"
)

type Query struct {
	queue []string
	start string
	Store Store
}

func (q *Query) V(start string) *Query {
	q.start = start
	return q
}

func (q *Query) Out(predicate string) *Query {
	q.queue = append(q.queue, predicate)
	return q
}

func (q *Query) Exec() []string {
	current := []string{q.start}
	for _, p := range q.queue {
		output := make(chan string)
		var wg sync.WaitGroup
		for _, v := range current {
			wg.Add(1)
			go func(subject string, predicate string) {
				for _, n := range q.Store.Objects(subject, predicate) {
					output <- n
				}
				wg.Done()
			}(v, p)
		}
		go func() {
			wg.Wait()
			close(output)
		}()
		current = make([]string, 0)
		for n := range output {
			current = append(current, n)
		}
	}
	return current
}

type Vertex struct {
	Value string
}
