package env_test

import (
	"testing"

	"github.com/aitorfernandez/roshambo/pkg/env"
)

func TestGet(t *testing.T) {
	env.Set("foo", "bar")

	if got, _ := env.Get("foo"); got != "bar" {
		t.Errorf("got %v want bar", got)
	}
}

func TestHget(t *testing.T) {
	env.Hset("hash", "baz", "qux")

	if got, _ := env.Hget("hash", "baz"); got != "qux" {
		t.Errorf("got %v want test", got)
	}
}
