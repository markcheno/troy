package troy

import ()

var store Store

func Init(s Store) {
	store = s
}

func Get(start string) *Query {
	q := new(Query)
	q.Store = store
	q.Vertices = []string{start}
	return q
}

func Update(start string) *Write {
	q := new(Write)
	q.Store = store
	return q.V(start)
}
