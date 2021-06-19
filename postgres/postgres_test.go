package postgres

import (
	"testing"
)

func TestGostgresConn(t *testing.T) {
	connString := ConnString("192.168.162.130", 5432, "toor333666", "postgres", "go_graphql_db")
	db, err := New(connString)
	if err != nil {
		t.Fatal(err)
	}

	if db.Ping() != nil {
		t.Fatal("invalid db")
	}
}
