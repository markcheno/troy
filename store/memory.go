package store

import (
//"log"
)

type Memory struct {
}

var cache map[string]map[string]map[string]bool

func (s *Memory) Create() {
	cache = make(map[string]map[string]map[string]bool)
}

func (s *Memory) Objects(subject string, predicate string) []string {
	var result []string
	for k := range cache[subject][predicate] {
		result = append(result, k)
	}
	return result
}

func (s *Memory) Triples(subject string) (string, string, string) {
	return "", "", ""
}

func (*Memory) Update(subject string, predicate string, object string) {
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

func (*Memory) Exists(subject string, predicate string, object string) bool {
	_, ok := cache[subject][predicate][object]
	return ok
}
