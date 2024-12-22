package ctxkey_test

import (
	"testing"

	"github.com/aitorfernandez/roshambo/pkg/ctxkey"
)

func TestString(t *testing.T) {
	var (
		str = "test"
		k   = ctxkey.New(str)
	)
	if got := k.String(); got != str {
		t.Errorf("want %s, got %v", str, got)
	}
}
