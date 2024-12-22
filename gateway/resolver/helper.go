package resolver

import (
	"fmt"

	"github.com/graph-gophers/graphql-go"
)

func gqlIDToString(id graphql.ID) string {
	return fmt.Sprint(id)
}
