package model

import (
	pb "github.com/aitorfernandez/roshambo/proto"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// AccountPrefix for prefix ids.
const AccountPrefix = "account"

// Account determines the account structure.
type Account struct {
	ID            string
	CreatedAt     int64
	Email         string
	LastRequestAt int64
	UpdatedAt     int64
}

// Scan scans an account row.
func (a *Account) Scan(scanner interface {
	Scan(dest ...interface{}) error
}) error {
	return scanner.Scan(
		&a.ID,
		&a.CreatedAt,
		&a.Email,
		&a.LastRequestAt,
		&a.UpdatedAt,
	)
}

// Validate validates fields for account model.
func (a Account) Validate() error {
	return validation.ValidateStruct(
		&a,
		validation.Field(&a.ID, validation.Required),
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.LastRequestAt, validation.Required),
	)
}

// Proto creates an account proto model from account.
func (a Account) Proto() *pb.Account {
	return &pb.Account{
		ID:    a.ID,
		Email: a.Email,
	}
}
