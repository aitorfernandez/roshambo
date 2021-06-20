package schema

import (
	"io/ioutil"
	"path/filepath"

	"github.com/aitorfernandez/roshambo/pkg/projectpath"
)

// String returns the GraphQL schema as string.
func String() (string, error) {
	var (
		content []byte
		err     error
		path    string
	)
	path = filepath.Join(projectpath.Base(), "gateway", "schema", "schema.graphql")
	if content, err = ioutil.ReadFile(path); err != nil {
		return "", err
	}
	return string(content), nil
}
