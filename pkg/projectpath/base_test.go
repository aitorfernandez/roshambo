package projectpath_test

import (
	"strings"
	"testing"

	"github.com/aitorfernandez/roshambo/pkg/projectpath"
)

func TestBase(t *testing.T) {
	if got := projectpath.Base(); strings.Index(got, "roshambo") < 1 {
		t.Errorf("wrong project path %v", got)
	}
}
