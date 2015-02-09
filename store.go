package troy

type Store interface {
	Objects(subject string, predicate string) []string
	Subjects(object string, predicate string) []string
	Update(subject string, predicate string, object string)
	Exists(subject string, predicate string, object string) bool
}
