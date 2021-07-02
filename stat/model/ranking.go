package model

import (
	pb "github.com/aitorfernandez/roshambo/proto"
)

// Ranking determines the stat structure.
type Ranking struct {
	ID          int32
	AccountID   string
	Draw        int32
	Lose        int32
	TotalRounds int32
	Win         int32
}

// Scan scans a ranking row.
func (r *Ranking) Scan(scanner interface {
	Scan(dest ...interface{}) error
}) error {
	return scanner.Scan(
		&r.ID,
		&r.AccountID,
		&r.Draw,
		&r.Lose,
		&r.TotalRounds,
		&r.Win,
	)
}

// Proto creates a ranking proto model from ranking.
func (r Ranking) Proto() *pb.Ranking {
	return &pb.Ranking{
		ID:          r.ID,
		AccountID:   r.AccountID,
		Draw:        r.Draw,
		Lose:        r.Lose,
		TotalRounds: r.TotalRounds,
		Win:         r.Win,
	}
}
