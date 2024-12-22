package resolver

import (
	"context"

	"github.com/aitorfernandez/roshambo/gateway/loader"
	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/graph-gophers/graphql-go"
)

func statRes(p *pb.Stat, err error) (*StatResolver, error) {
	if err != nil {
		return nil, err
	}
	return &StatResolver{p}, nil
}

func newStatsByAccount(ctx context.Context, accountID string) ([]*StatResolver, error) {
	stats, err := loader.LoadStatsByAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}
	ss := make([]*StatResolver, 0, len(stats))
	for _, s := range stats {
		ss = append(ss, &StatResolver{s})
	}
	return ss, nil
}

// StatResolver resolves Stat GraphQL type.
type StatResolver struct {
	stat *pb.Stat
}

// ID resolves ID field.
func (r StatResolver) ID() graphql.ID {
	return graphql.ID(r.stat.ID)
}

// ComputerMove resolves computerMove field.
func (r StatResolver) ComputerMove() int32 {
	return r.stat.ComputerMove
}

// PlayerMove resolves playerMove field.
func (r StatResolver) PlayerMove() int32 {
	return r.stat.PlayerMove
}

// Round resolves round field.
func (r StatResolver) Round() int32 {
	return r.stat.Round
}
