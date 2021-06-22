package client

import (
	"context"

	pb "github.com/aitorfernandez/roshambo/proto"
)

// GetProfileByAccount sends GetProfileByAccount request to profileServer.
func (c Client) GetProfileByAccount(ctx context.Context, req *pb.GetProfileByAccountReq) (*pb.Profile, error) {
	return pb.NewProfileServiceClient(c.profileConn).GetProfileByAccount(ctx, req)
}
