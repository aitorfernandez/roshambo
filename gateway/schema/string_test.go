package schema_test

import (
	"testing"

	"github.com/aitorfernandez/roshambo/gateway/schema"
)

func TestString(t *testing.T) {
	if got, _ := schema.String(); got == "" {
		t.Error("invalid empty schema")
	}
}
