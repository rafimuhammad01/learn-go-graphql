package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)


type Db struct {
	*sql.DB
}

func NewConn(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// Check that our connection is good
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

func ConnString(host string, port string, user string, password string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)
}
