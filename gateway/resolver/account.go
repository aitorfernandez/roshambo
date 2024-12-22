package resolver

import (
	"context"

	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/graph-gophers/graphql-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func accountRes(a *pb.Account, err error) (*AccountResolver, error) {
	if err != nil {
		return nil, err
	}
	return &AccountResolver{a}, nil
}

// AccountResolver resolves Account GraphQL type.
type AccountResolver struct {
	account *pb.Account
}

// ID resolves id field.
func (r AccountResolver) ID() graphql.ID {
	return graphql.ID(r.account.ID)
}

// Email resolves email field.
func (r AccountResolver) Email() string {
	return r.account.Email
}

// Profile resolves profile field.
func (r AccountResolver) Profile(ctx context.Context) (*ProfileResolver, error) {
	p, err := newProfileByAccount(ctx, r.account.ID)
	if err != nil {
		if e, ok := status.FromError(err); ok && e.Code() == codes.NotFound {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

// Stats resolves stats field.
func (r AccountResolver) Stats(ctx context.Context) ([]*StatResolver, error) {
	return newStatsByAccount(ctx, r.account.ID)
}
