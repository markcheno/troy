package troy

import (
	"github.com/ironbay/troy/cassandra"
)

var store Store

func init() {
	store = new(cassandra.Store)
	store.Create()
}

func Get(start string) *Query {
	q := new(Query)
	q.Store = store
	return q.V(start)
}

func Update(start string) *Write {
	q := new(Write)
	q.Store = store
	return q.V(start)
}
