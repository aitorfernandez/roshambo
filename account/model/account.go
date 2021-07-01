package model

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aitorfernandez/roshambo/pkg/env"
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

// GenerateToken generates a token.
func (a Account) GenerateToken(ts int64, ip string) (string, error) {
	var (
		err    error
		secret string
		token  = fmt.Sprintf("%s%d%d%s", a.ID, a.LastRequestAt, ts, ip)
	)
	if secret, err = env.Hget("app", "secret"); err != nil {
		return "", err
	}
	return strings.Join([]string{
		strconv.FormatInt(ts, 10),
		hex.EncodeToString(hmacsha256([]byte(secret), []byte(token))),
	}, "-"), nil
}

// ValidateToken validates a token with an account.
func (a Account) ValidateToken(token, ip string) bool {
	var (
		err error
		t   string
		ts  int64
	)
	if ts, err = strconv.ParseInt(strings.Split(token, "-")[0], 10, 64); err != nil {
		return false
	}
	if time.Now().Unix() > (ts + (15 * 60)) {
		return false
	}
	if t, err = a.GenerateToken(ts, ip); err != nil {
		return false
	}
	return t == token
}

func hmacsha256(key []byte, data []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(data)
	return hash.Sum(nil)
}
