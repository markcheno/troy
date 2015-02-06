package memory

import (
//"log"
)

type Store struct {
}

var cache map[string]map[string]map[string]bool

func (s *Store) Create() {
	cache = make(map[string]map[string]map[string]bool)
}

func (s *Store) Objects(subject string, predicate string) []string {
	var result []string
	for k := range cache[subject][predicate] {
		result = append(result, k)
	}
	return result
}

func (s *Store) Triples(subject string) (string, string, string) {
	return "", "", ""
}

func (*Store) Update(subject string, predicate string, object string) {
	_, ok := cache[subject]
	if !ok {
		cache[subject] = make(map[string]map[string]bool)
	}
	_, ok = cache[subject][predicate]
	if !ok {
		o := make(map[string]bool)
		cache[subject][predicate] = o
	}
	cache[subject][predicate][object] = true
}

func (*Store) Exists(subject string, predicate string, object string) bool {
	_, ok := cache[subject][predicate][object]
	return ok
}
