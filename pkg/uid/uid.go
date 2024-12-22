package uid

import (
	"strings"

	"github.com/lucsky/cuid"
)

// New returns a new uid with prefix value.
func New(str string) string {
	return strings.Join([]string{
		str[:3],
		cuid.New(),
	}, "_")
}
