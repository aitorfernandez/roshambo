package resolver

import (
	"context"
	"fmt"

	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/graph-gophers/graphql-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func rankingRes(r *pb.Ranking, err error) (*RankingResolver, error) {
	if err != nil {
		return nil, err
	}
	return &RankingResolver{r}, nil
}

// RankingResolver resolves ranking GraphQL type.
type RankingResolver struct {
	ranking *pb.Ranking
}

// ID resolves ID field.
func (r RankingResolver) ID() graphql.ID {
	return graphql.ID(fmt.Sprint(r.ranking.ID))
}

// Draw resolves draw field.
func (r RankingResolver) Draw() int32 {
	return r.ranking.Draw
}

// Lose resolves lose field.
func (r RankingResolver) Lose() int32 {
	return r.ranking.Lose
}

// TotalRounds resolves totalRounds field.
func (r RankingResolver) TotalRounds() int32 {
	return r.ranking.TotalRounds
}

// Username resolves username field.
func (r RankingResolver) Username(ctx context.Context) (string, error) {
	p, err := newProfileByAccount(ctx, r.ranking.AccountID)
	if err != nil {
		if e, ok := status.FromError(err); ok && e.Code() == codes.NotFound {
			return "anonymous", nil
		}
		return "", err
	}
	return p.Username(), nil
}

// Win resolves win field.
func (r RankingResolver) Win() int32 {
	return r.ranking.Win
}
