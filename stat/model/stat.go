package model

import (
	pb "github.com/aitorfernandez/roshambo/proto"
	validation "github.com/go-ozzo/ozzo-validation"
)

// Stat determines the stat structure.
type Stat struct {
	ID           string
	AccountID    string
	ComputerMove int32
	PlayerMove   int32
	Round        int32
}

// Scan scans a stat row.
func (s *Stat) Scan(scanner interface {
	Scan(dest ...interface{}) error
}) error {
	return scanner.Scan(
		&s.ID,
		&s.AccountID,
		&s.ComputerMove,
		&s.PlayerMove,
		&s.Round,
	)
}

// Validate validates fields for profile model.
func (s Stat) Validate() error {
	return validation.ValidateStruct(
		&s,
		validation.Field(&s.AccountID, validation.Required),
		validation.Field(&s.ComputerMove, validation.Required),
		validation.Field(&s.PlayerMove, validation.Required),
	)
}

// Proto creates a stat proto model from stat.
func (s Stat) Proto() *pb.Stat {
	return &pb.Stat{
		ID:           s.ID,
		AccountID:    s.AccountID,
		ComputerMove: s.ComputerMove,
		PlayerMove:   s.PlayerMove,
		Round:        s.Round,
	}
}
