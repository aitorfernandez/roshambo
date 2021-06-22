package store

import (
	"context"
	"database/sql"

	"github.com/aitorfernandez/roshambo/pkg/postgres"
	"github.com/aitorfernandez/roshambo/profile/model"
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

func profileRow(row *sql.Row) (*model.Profile, error) {
	var (
		err error
		p   model.Profile
	)
	if err = p.Scan(row); err != nil {
		return nil, err
	}
	return &p, nil
}

// ProfileByAccount gets a profile for the given accountID.
func (s Store) ProfileByAccount(ctx context.Context, accountID string) (*model.Profile, error) {
	q := `
	select * from profile where account_id = $1
	`
	row := s.db.QueryRowContext(ctx, q, accountID)
	return profileRow(row)
}
