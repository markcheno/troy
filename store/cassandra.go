package store

import (
	"github.com/gocql/gocql"
)

var cassandra *gocql.Session

type Cassandra struct{}

func (*Cassandra) Create(host string, keyspace string) {
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = keyspace
	var err error
	cassandra, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	err = cassandra.Query("CREATE TABLE IF NOT EXISTS subjects ( subject text, predicate text, object text, updated timestamp, PRIMARY KEY(subject, predicate, object) )").Exec()
	err = cassandra.Query("CREATE TABLE IF NOT EXISTS objects ( object text, predicate text, subject text, updated timestamp, PRIMARY KEY(object, predicate, subject) )").Exec()
	if err != nil {
		panic(err)
	}
}

func (*Cassandra) Objects(subject string, predicate string) []string {
	i := cassandra.Query("SELECT object FROM subjects WHERE subject = ? AND predicate = ?", subject, predicate).Iter()
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

func (*Cassandra) Subjects(object string, predicate string) []string {
	i := cassandra.Query("SELECT subject FROM objects WHERE object= ? AND predicate = ?", object, predicate).Iter()
	result := make([]string, 0)
	var subject string
	for i.Scan(&subject) {
		result = append(result, subject)
	}
	err := i.Close()
	if err != nil {
		panic(err)
	}
	return result
}

func (*Cassandra) Update(subject string, predicate string, object string) {
	err := cassandra.Query("UPDATE subjects SET updated = dateof(now()) WHERE object = ? AND subject = ? AND predicate = ?", object, subject, predicate).Exec()
	err = cassandra.Query("UPDATE objects SET updated = dateof(now()) WHERE object = ? AND subject = ? AND predicate = ?", object, subject, predicate).Exec()
	if err != nil {
		panic(err)
	}
}

func (*Cassandra) Exists(subject string, predicate string, object string) bool {
	m, err := cassandra.Query("SELECT * FROM triples WHERE object = ? AND subject = ? AND predicate = ?", object, subject, predicate).Iter().SliceMap()
	if err != nil {
		panic(err)
	}
	return len(m) > 0
}
