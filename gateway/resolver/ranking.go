package resolver

import (
	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/graph-gophers/graphql-go"
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

// ID resolves id field.
func (r RankingResolver) ID() graphql.ID {
	return graphql.ID(r.ranking.ID)
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

// Win resolves win field.
func (r RankingResolver) Win() int32 {
	return r.ranking.Win
}
