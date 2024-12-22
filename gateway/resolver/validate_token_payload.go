package resolver

import (
	"github.com/aitorfernandez/roshambo/gateway/jwt"
	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/graph-gophers/graphql-go"
)

func validateTokenPayloadRes(a *pb.Account, err error) (*ValidateTokenPayloadResolver, error) {
	if err != nil {
		return nil, err
	}
	return &ValidateTokenPayloadResolver{a}, nil
}

// ValidateTokenPayloadResolver resolves ValidateTokenPayload GraphQL type.
type ValidateTokenPayloadResolver struct {
	account *pb.Account
}

// ID resolves id field.
func (r ValidateTokenPayloadResolver) ID() graphql.ID {
	return graphql.ID(r.account.ID)
}

// JWT resolves jwt field.
func (r ValidateTokenPayloadResolver) JWT() (string, error) {
	return jwt.Make(r.account.ID)
}
