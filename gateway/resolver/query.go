package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

type accountArgs struct {
	ID graphql.ID
}

// Account resolves account query.
func (r Resolver) Account(ctx context.Context, args accountArgs) (*AccountResolver, error) {
	return &AccountResolver{}, nil
}
