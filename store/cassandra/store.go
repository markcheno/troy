package cassandra

import (
	"github.com/gocql/gocql"
)

var cassandra *gocql.Session

type Store struct{}

func (*Store) Create(host string, keyspace string) {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = keyspace
	var err error
	cassandra, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	err = cassandra.Query("CREATE TABLE IF NOT EXISTS triples ( subject text, predicate text, object text, updated timestamp, PRIMARY KEY(subject, predicate, object) )").Exec()
	if err != nil {
		panic(err)
	}
}

func (*Store) Objects(subject string, predicate string) []string {
	i := cassandra.Query("SELECT object FROM triples WHERE subject = ? AND predicate = ?", subject, predicate).Iter()
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
	err := cassandra.Query("UPDATE triples SET updated = dateof(now()) WHERE object = ? AND subject = ? AND predicate = ?", object, subject, predicate).Exec()
	if err != nil {
		panic(err)
	}
}

func (*Store) Exists(subject string, predicate string, object string) bool {
	m, err := cassandra.Query("SELECT * FROM triples WHERE object = ? AND subject = ? AND predicate = ?", object, subject, predicate).Iter().SliceMap()
	if err != nil {
		panic(err)
	}
	return len(m) > 0
}
