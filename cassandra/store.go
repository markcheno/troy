package cassandra

import (
	"github.com/gocql/gocql"
)

var cassandra *gocql.Session

type Store struct{}

func (*Store) Create() {
	cluster := gocql.NewCluster("localhost")
	var err error
	cassandra, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	err = cassandra.Query("CREATE TABLE IF NOT EXISTS troy.triples ( subject text, predicate text, object text, updated timestamp, PRIMARY KEY(subject, predicate, object) )").Exec()
	if err != nil {
		panic(err)
	}
}

func (*Store) Objects(subject string, predicate string) []string {
	i := cassandra.Query("SELECT object FROM troy.triples WHERE subject = ? AND predicate = ?", subject, predicate).Iter()
	result := make([]string, 0)
	var object string
	for i.Scan(&object) {
		result = append(result, object)
	}
	err := i.Close()
	if err != nil {
		panic(err)
	}
	return result
}

func (*Store) Triples(subject string) (string, string, string) {
	return "subject", "predicate", "object"
}

func (*Store) Update(subject string, predicate string, object string) {
	err := cassandra.Query("UPDATE troy.triples SET updated = dateof(now()) WHERE object = ? AND subject = ? AND predicate = ?", object, subject, predicate).Exec()
	if err != nil {
		panic(err)
	}
}
