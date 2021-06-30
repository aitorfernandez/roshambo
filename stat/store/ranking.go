package store

import (
	"context"
	"database/sql"

	"github.com/aitorfernandez/roshambo/stat/model"
)

func rankingRows(rows *sql.Rows, err error) ([]*model.Ranking, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rr := []*model.Ranking{}
	for rows.Next() {
		var (
			err error
			r   model.Ranking
		)
		if err = r.Scan(rows); err != nil {
			return nil, err
		}
		rr = append(rr, &r)
	}
	return rr, nil
}

// Rankings select all rankings.
func (s Store) Rankings(ctx context.Context) ([]*model.Ranking, error) {
	q := `
	select * from ranking
	`
	rows, err := s.db.QueryContext(ctx, q)
	return rankingRows(rows, err)
}
