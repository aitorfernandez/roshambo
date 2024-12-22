package resolver

import (
	"context"

	"github.com/aitorfernandez/roshambo/gateway/loader"
	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/graph-gophers/graphql-go"
)

func profileRes(p *pb.Profile, err error) (*ProfileResolver, error) {
	if err != nil {
		return nil, err
	}
	return &ProfileResolver{p}, nil
}

func newProfileByAccount(ctx context.Context, accountID string) (*ProfileResolver, error) {
	return profileRes(loader.LoadProfileByAccount(ctx, accountID))
}

// ProfileResolver resolves Profile GraphQL type.
type ProfileResolver struct {
	profile *pb.Profile
}

// AccountID resolves accountID field.
func (r ProfileResolver) AccountID() graphql.ID {
	return graphql.ID(r.profile.AccountID)
}

// Avatar resolves avatar field.
func (r ProfileResolver) Avatar() *string {
	return r.profile.Avatar
}

// Username resolves username field.
func (r ProfileResolver) Username() string {
	return r.profile.Username
}
