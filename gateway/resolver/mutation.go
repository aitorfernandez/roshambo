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
	return validateTokenPayloadRes(res, err)
}

type createStatInput struct {
	AccountID  graphql.ID
	PlayerMove int32
}

type createStatArgs struct {
	Input createStatInput
}

// CreateStat resolves createStat mutation.
func (r Resolver) CreateStat(ctx context.Context, args createStatArgs) (*StatResolver, error) {
	res, err := r.client.CreateStat(ctx, &pb.CreateStatReq{
		Stat: &pb.Stat{
			AccountID:  gqlIDToString(args.Input.AccountID),
			PlayerMove: args.Input.PlayerMove,
		},
	})
	return statRes(res, err)
}

type deleteStatsByAccountArgs struct {
	AccountID graphql.ID
}

// DeleteStatsByAccount relsolves deleteStatsByAccount mutation.
func (r Resolver) DeleteStatsByAccount(ctx context.Context, args deleteStatsByAccountArgs) (int32, error) {
	res, err := r.client.DeleteStatsByAccount(ctx, &pb.DeleteStatsByAccountReq{
		AccountID: gqlIDToString(args.AccountID),
	})
	if err != nil {
		return 0, nil
	}
	return res, nil
}

//
// Profile
//

type setProfileInput struct {
	AccountID graphql.ID
	Avatar    *string
	Username  string
}

type setProfileArgs struct {
	Input setProfileInput
}

// SetProfile resolves setProfile mutation.
func (r Resolver) SetProfile(ctx context.Context, args setProfileArgs) (*ProfileResolver, error) {
	res, err := r.client.SetProfile(ctx, &pb.SetProfileReq{
		Profile: &pb.Profile{
			AccountID: gqlIDToString(args.Input.AccountID),
			Avatar:    args.Input.Avatar,
			Username:  args.Input.Username,
		},
	})
	return profileRes(res, err)
}
