package model_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/aitorfernandez/roshambo/account/model"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	t.Parallel()

	a := &model.Account{
		ID:            "1d",
		LastRequestAt: int64(1257894000),
	}
	token, err := a.GenerateToken(a.LastRequestAt, "0.0.0.0")
	assert.NoError(t, err)
	assert.Equal(t, token, "1257894000-5316bc405bf002a1df3b28b901111ab669a023397aa89edbca8a2a34a46ce91a")
}

func TestCompareToken(t *testing.T) {
	t.Parallel()

	a := &model.Account{
		ID:            "1d",
		LastRequestAt: time.Now().Unix(),
	}

	t.Run("valid token", func(t *testing.T) {
		token, _ := a.GenerateToken(a.LastRequestAt, "0.0.0.0")
		ok := a.ValidateToken(token, "0.0.0.0")
		assert.True(t, ok)
	})

	t.Run("malformed token", func(t *testing.T) {
		token, _ := a.GenerateToken(a.LastRequestAt, "0.0.0.0")
		token = strings.ReplaceAll(
			token,
			"-",
			"X",
		)
		ok := a.ValidateToken(token, "0.0.0.0")
		assert.False(t, ok)
	})

	t.Run("invalid ip", func(t *testing.T) {
		token, _ := a.GenerateToken(a.LastRequestAt, "0.0.0.0")
		ok := a.ValidateToken(token, "1.1.1.1")
		assert.False(t, ok)
	})

	t.Run("expired token", func(t *testing.T) {
		token, _ := a.GenerateToken(a.LastRequestAt, "0.0.0.0")
		token = strings.ReplaceAll(
			token,
			fmt.Sprintf("%d%s", a.LastRequestAt, "-"),
			fmt.Sprintf("%d%s", a.LastRequestAt-(20*60), "-"),
		)
		ok := a.ValidateToken(token, "0.0.0.0")
		assert.False(t, ok)
	})
}

func TestProto(t *testing.T) {
	t.Parallel()

	var a model.Account
	faker.FakeData(&a)

	pb := a.Proto()
	assert.NotNil(t, pb.ID)
	assert.NotNil(t, pb.Email)
}

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		a := &model.Account{
			ID:            "id",
			Email:         "test@test.com",
			LastRequestAt: time.Now().Unix(),
		}
		err := a.Validate()
		assert.NoError(t, err)
	})

	t.Run("missing id", func(t *testing.T) {
		a := &model.Account{
			Email:         "test@test.com",
			LastRequestAt: time.Now().Unix(),
		}
		err := a.Validate()
		assert.NotNil(t, err)
	})

	t.Run("missing email", func(t *testing.T) {
		a := &model.Account{
			ID:            "id",
			LastRequestAt: time.Now().Unix(),
		}
		err := a.Validate()
		assert.NotNil(t, err)
	})

	t.Run("missing lastRequestAt", func(t *testing.T) {
		a := &model.Account{
			ID:    "id",
			Email: "test@test.com",
		}
		err := a.Validate()
		assert.NotNil(t, err)
	})
}
