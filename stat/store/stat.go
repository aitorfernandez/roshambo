package store

import (
	"context"
	"database/sql"

	"github.com/aitorfernandez/roshambo/stat/model"
)

func statRow(row *sql.Row) (*model.Stat, error) {
	var (
		err error
		s   model.Stat
	)
	if err = s.Scan(row); err != nil {
		return nil, err
	}
	return &s, nil
}

func statRows(rows *sql.Rows, err error) ([]*model.Stat, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ss := []*model.Stat{}
	for rows.Next() {
		var (
			err error
			s   model.Stat
		)
		if err = s.Scan(rows); err != nil {
			return nil, err
		}
		ss = append(ss, &s)
	}
	return ss, nil
}

// CreateStat creates a new stat.
func (s Store) CreateStat(ctx context.Context, st *model.Stat) (*model.Stat, error) {
	q := `
	with m as (
	  select coalesce(max(round) + 1, 1) from stat where account_id = $1
	)
	insert into stat (account_id, computer_move, player_move, round)
	values
	  ($1, $2, $3, (select * from m))
	returning
	  *
	`
	row := s.db.QueryRowContext(ctx, q, st.AccountID, st.ComputerMove, st.PlayerMove)
	return statRow(row)
}

// DeleteStatsByAccount delete all stats for the given account.
func (s Store) DeleteStatsByAccount(ctx context.Context, accountID string) (int32, error) {
	q := `
	with d as (
	  delete from stat
	  where
		account_id = $1
	  returning
		*
	)
	select count(*) from d
	`
	var (
		count int32
		err   error
	)
	if err = s.db.QueryRowContext(ctx, q, accountID).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// StatsByAccount selects all stats for the given account.
func (s Store) StatsByAccount(ctx context.Context, accountID string) ([]*model.Stat, error) {
	q := `
	select * from stat where account_id = $1 order by round desc
	`
	rows, err := s.db.QueryContext(ctx, q, accountID)
	return statRows(rows, err)
}
