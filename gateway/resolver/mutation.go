package resolver

import (
	"context"

	"github.com/aitorfernandez/roshambo/gateway/middleware"
	pb "github.com/aitorfernandez/roshambo/proto"
)

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
