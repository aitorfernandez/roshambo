package store_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aitorfernandez/roshambo/account/model"
	"github.com/aitorfernandez/roshambo/account/store"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

var (
	columns = []string{"id", "created_at", "email", "last_request_at", "updated_at"}
	rows    = sqlmock.NewRows(columns)
)

func TestAccount(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	a := fakeAccount()
	s := &store.Store{db}
	q := `
	select * from account where id = $1
	`
	mock.ExpectQuery(regexp.QuoteMeta(q)).
		WithArgs(a.ID).
		WillReturnRows(rows.AddRow(a.ID, a.CreatedAt, a.Email, a.LastRequestAt, a.UpdatedAt))

	_, err = s.Account(context.Background(), a.ID)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "there were unfulfilled expectations: %s", err)
}

func TestSetAccount(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	a := fakeAccount()
	s := &store.Store{db}
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
	mock.ExpectQuery(regexp.QuoteMeta(q)).
		WithArgs(a.ID, a.Email, a.LastRequestAt).
		WillReturnRows(rows.AddRow(a.ID, a.CreatedAt, a.Email, a.LastRequestAt, a.UpdatedAt))

	_, err = s.SetAccount(context.Background(), a)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "there were unfulfilled expectations: %s", err)
}

func TestUpdateLastRequestAt(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	a := fakeAccount()
	s := &store.Store{db}
	q := `
	update account
	set
	  last_request_at = $1
	where
	  id = $2
	returning
	  *
	`
	mock.ExpectQuery(regexp.QuoteMeta(q)).
		WithArgs(a.LastRequestAt, a.ID).
		WillReturnRows(rows.AddRow(a.ID, a.CreatedAt, a.Email, a.LastRequestAt, a.UpdatedAt))

	_, err = s.UpdateLastRequestAt(context.Background(), a.ID, a.LastRequestAt)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "there were unfulfilled expectations: %s", err)
}

func fakeAccount() *model.Account {
	var a model.Account
	faker.FakeData(&a)
	return &a
}
