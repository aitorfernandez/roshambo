package uid_test

import (
	"strings"
	"testing"

	"github.com/aitorfernandez/roshambo/pkg/uid"
)

func TestNew(t *testing.T) {
	got := strings.Split(uid.New("teacher"), "_")[0]
	if got != "acc" {
		t.Errorf("want a correct prefix, got %s", got)
	}
}
