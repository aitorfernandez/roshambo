package resolver

import (
	"context"

	"github.com/aitorfernandez/roshambo/gateway/middleware"
	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/graph-gophers/graphql-go"
)

//
// Account
//

type getStartedArgs struct {
	Email string
}

// GetStarted resolves getStarted mutation.
func (r Resolver) GetStarted(ctx context.Context, args getStartedArgs) (bool, error) {
	res, err := r.client.GetStarted(ctx, &pb.GetStartedReq{
		Email: args.Email,
		IP:    ctx.Value(middleware.RemoteAddrKey).(string),
	})
	if err != nil {
		return false, err
	}
	return res.Value, nil
}

type validateTokenInput struct {
	ID    graphql.ID
	Token string
}

type validateTokenArgs struct {
	Input validateTokenInput
}

// ValidateToken resolves validateToken mutation.
func (r Resolver) ValidateToken(ctx context.Context, args validateTokenArgs) (*ValidateTokenPayloadResolver, error) {
	res, err := r.client.ValidateToken(ctx, &pb.ValidateTokenReq{
		ID:    gqlIDToString(args.Input.ID),
		IP:    ctx.Value(middleware.RemoteAddrKey).(string),
		Token: args.Input.Token,
	})
	return validateTokenRes(res, err)
}
