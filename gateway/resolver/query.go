package resolver

import (
	"context"

	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/graph-gophers/graphql-go"
)

type accountArgs struct {
	ID graphql.ID
}

// Account resolves account query.
func (r Resolver) Account(ctx context.Context, args accountArgs) (*AccountResolver, error) {
	res, err := r.client.GetAccount(ctx, &pb.GetAccountReq{
		ID: gqlIDToString(args.ID),
	})
	return accountRes(res, err)
}
