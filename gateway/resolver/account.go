package resolver

import (
	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/graph-gophers/graphql-go"
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
