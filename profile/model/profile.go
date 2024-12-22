package model

import (
	"database/sql"

	pb "github.com/aitorfernandez/roshambo/proto"
	validation "github.com/go-ozzo/ozzo-validation"
)

// ProfilePrefix for id generation.
const ProfilePrefix = "profile"

// Profile determines the profile structure.
type Profile struct {
	ID        string
	AccountID string
	Avatar    *string
	CreatedAt int64
	UpdatedAt int64
	Username  string
}

// Scan scans a profile row.
func (p *Profile) Scan(scanner interface {
	Scan(dest ...interface{}) error
}) error {
	var (
		err    error
		avatar sql.NullString
	)
	if err = scanner.Scan(
		&p.ID,
		&p.AccountID,
		&avatar,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Username,
	); err != nil {
		return err
	}
	if avatar.Valid {
		p.Avatar = &avatar.String
	}
	return nil
}

// Validate validates fields for profile model.
func (p Profile) Validate() error {
	return validation.ValidateStruct(
		&p,
		validation.Field(&p.ID, validation.Required),
		validation.Field(&p.AccountID, validation.Required),
	)
}

// Proto creates a profile proto model from profile.
func (p Profile) Proto() *pb.Profile {
	return &pb.Profile{
		ID:        p.ID,
		AccountID: p.AccountID,
		Avatar:    p.Avatar,
		Username:  p.Username,
	}
}
