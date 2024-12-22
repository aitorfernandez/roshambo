package store

import (
	"context"
	"database/sql"

	"github.com/aitorfernandez/roshambo/account/model"
	"github.com/aitorfernandez/roshambo/pkg/postgres"
)

// Store holds the db connection and db data access layer.
type Store struct {
	DB *sql.DB
}

// New creates a new Data.
func New(dataSource string) (*Store, error) {
	db, err := postgres.Conn(dataSource)
	if err != nil {
		return nil, err
	}
	return &Store{db}, nil
}

func accountRow(row *sql.Row) (*model.Account, error) {
	var (
		a   model.Account
		err error
	)
	if err = a.Scan(row); err != nil {
		return nil, err
	}
	return &a, nil
}

// Account selects an account for the given id.
func (s Store) Account(ctx context.Context, id string) (*model.Account, error) {
	q := `
	select * from account where id = $1
	`
	row := s.DB.QueryRowContext(ctx, q, id)
	return accountRow(row)
}

// SetAccount inserts or updates an existing account.
func (s Store) SetAccount(ctx context.Context, a *model.Account) (*model.Account, error) {
	q := `
	insert into account (id, email, last_request_at)
	values
	  ($1, $2, $3)
	on conflict
	  (email)
	do update set
	  last_request_at = $3
	returning
	  *
	`
	row := s.DB.QueryRowContext(
		ctx, q, a.ID, a.Email, a.LastRequestAt,
	)
	return accountRow(row)
}

// UpdateLastRequestAt updates last_request_at for the given account.
func (s Store) UpdateLastRequestAt(ctx context.Context, id string, requestAt int64) (*model.Account, error) {
	q := `
	update account
	set
	  last_request_at = $1
	where
	  id = $2
	returning
	  *
	`
	row := s.DB.QueryRowContext(ctx, q, requestAt, id)
	return accountRow(row)
}
