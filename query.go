package troy

import (
	"sync"
)

type Query struct {
	Vertices  []string
	predicate string
	Store     Store
}

func (q *Query) All() *Query {
	output := make(chan string)
	var wg sync.WaitGroup
	for _, vertex := range q.Vertices {
		wg.Add(1)
		go func(subject string, predicate string) {
			for _, n := range q.Store.Objects(subject, predicate) {
				output <- n
			}
			wg.Done()
		}(vertex, q.predicate)
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	q.Vertices = make([]string, 0)
	for n := range output {
		q.Vertices = append(q.Vertices, n)
	}
	return q
}

func (q *Query) V(end string) *Query {
	success := true
	for _, vertex := range q.Vertices {
		if !q.Store.Exists(vertex, q.predicate, end) {
			success = false
			break
		}
	}
	q.Vertices = make([]string, 0)
	if success {
		q.Vertices = append(q.Vertices, end)
	}
	return q
}

func (q *Query) Out(predicate string) *Query {
	q.predicate = predicate
	return q
}
