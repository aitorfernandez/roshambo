package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq" // driver for Postgres
)

var (
	defaultDB *sql.DB
	once      sync.Once
)

func open(dataSource string) (*sql.DB, error) {
	var (
		DB  *sql.DB
		err error
	)
	if DB, err = sql.Open("postgres", dataSource); err != nil {
		return nil, fmt.Errorf("pkg.postgres.conn sql.Open %w", err)
	}
	if err = DB.Ping(); err != nil {
		return nil, fmt.Errorf("pkg.postgres.conn DB.Ping %w", err)
	}
	return DB, err
}

// Conn returns a connection to the database and returns a DB struct.
func Conn(dataSource string) (*sql.DB, error) {
	var err error
	once.Do(func() {
		defaultDB, err = open(dataSource)
	})
	return defaultDB, err
}
