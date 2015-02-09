package troy

import ()

var store Store

func Init(s Store) {
	store = s
}

func V(start ...string) *Query {
	q := new(Query)
	q.Store = store
	q.Result = start
	return q
}

func Update(start string) *Write {
	q := new(Write)
	q.Queue = []string{start}
	q.Store = store
	return q
}
