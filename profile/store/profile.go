package store

import (
	"context"
	"database/sql"

	"github.com/aitorfernandez/roshambo/pkg/null"
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

// SetProfile creates or updates the given profile.
func (s Store) SetProfile(ctx context.Context, p *model.Profile) (*model.Profile, error) {
	q := `
	insert into profile (id, account_id, avatar, username)
	values
	  ($1, $2, $3, $4)
	on conflict
	  (account_id)
	do update set
	  avatar = coalesce($3, profile.avatar),
	  username = $4
	returning
	  *
	`
	row := s.db.QueryRowContext(
		ctx, q, p.ID, p.AccountID, null.SQLString(p.Avatar), p.Username,
	)
	return profileRow(row)
}
