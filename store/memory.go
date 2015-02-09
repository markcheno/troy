package store

import (
//"log"
)

type Memory struct {
	subjects Cache
	objects  Cache
}

type Cache map[string]map[string]map[string]bool

func (s *Memory) Create() {
	s.subjects = Cache{}
	s.objects = Cache{}
}

func (s *Memory) Objects(subject string, predicate string) []string {
	var result []string
	for k := range s.subjects[subject][predicate] {
		result = append(result, k)
	}
	return result
}

func (s *Memory) Subjects(object string, predicate string) []string {
	var result []string
	for k := range s.objects[object][predicate] {
		result = append(result, k)
	}
	return result
}

func (s *Memory) Update(subject string, predicate string, object string) {
	s.updateCache(s.subjects, subject, predicate, object)
	s.updateCache(s.objects, object, predicate, subject)
}

func (*Memory) updateCache(cache Cache, first string, second string, third string) {
	_, ok := cache[first]
	if !ok {
		cache[first] = make(map[string]map[string]bool)
	}
	_, ok = cache[first][second]
	if !ok {
		cache[first][second] = make(map[string]bool)
	}
	cache[first][second][third] = true
}

func (s *Memory) Exists(subject string, predicate string, object string) bool {
	_, ok := s.subjects[subject][predicate][object]
	return ok
}
