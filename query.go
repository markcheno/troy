package troy

import (
	"sync"
)

type Query struct {
	Store        Store
	Instructions []Traversal
	Result       []string
}

type Traversal struct {
	Type  string
	Value string
}

func (q *Query) Has(predicate string, object string) *Query {
	valid := []string{}
	for _, vertex := range q.Result {
		if q.Store.Exists(vertex, predicate, object) {
			valid = append(valid, vertex)
		}
	}
	q.Result = valid
	return q
}

func (q *Query) Out(predicate string) *Query {
	q.edge(predicate, q.Store.Objects)
	return q
}

func (q *Query) In(predicate string) *Query {
	q.edge(predicate, q.Store.Subjects)
	return q
}

func (q *Query) Group() map[string]int {
	result := map[string]int{}
	for _, vertex := range q.Result {
		result[vertex]++
	}
	return result
}

func (q *Query) edge(predicate string, direction func(string, string) []string) []string {
	output := make(chan string)
	var wg sync.WaitGroup
	for _, vertex := range q.Result {
		wg.Add(1)
		go func(subject string) {
			for _, n := range direction(subject, predicate) {
				output <- n
			}
			wg.Done()
		}(vertex)
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	q.Result = make([]string, 0)
	for n := range output {
		q.Result = append(q.Result, n)
	}
	return q.Result
}
