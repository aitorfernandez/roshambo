package resolver

import (
	"context"

	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/graph-gophers/graphql-go"
	"google.golang.org/protobuf/types/known/emptypb"
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

// Rankings resolves rankings query.
func (r Resolver) Rankings(ctx context.Context) ([]*RankingResolver, error) {
	rankings, err := r.client.ListRankings(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	rr := make([]*RankingResolver, 0, len(rankings))
	for _, r := range rankings {
		rr = append(rr, &RankingResolver{r})
	}
	return rr, nil
}
