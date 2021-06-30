package store

import (
	"database/sql"

	"github.com/aitorfernandez/roshambo/pkg/postgres"
)

// Store holds the db connection and db data access layer.
type Store struct {
	db *sql.DB
}

// New creates a new Data.
func New(dataSource string) (*Store, error) {
	db, err := postgres.Conn(dataSource)
	if err != nil {
		return nil, err
	}
	return &Store{db}, nil
}
