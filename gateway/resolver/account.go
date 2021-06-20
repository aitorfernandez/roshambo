package resolver

import (
	"github.com/graph-gophers/graphql-go"
)

// AccountResolver resolves Account GraphQL type.
type AccountResolver struct{}

// ID resolves id field.
func (r AccountResolver) ID() graphql.ID {
	return graphql.ID("1d")
}

// Email resolves email field.
func (r AccountResolver) Email() string {
	return "foo@gmail.com"
}
