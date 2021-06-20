package client

import (
	"context"

	pb "github.com/aitorfernandez/roshambo/proto"
)

// GetAccount sends GetAccount request ro accountServer.
func (c Client) GetAccount(ctx context.Context, req *pb.GetAccountReq) (*pb.Account, error) {
	return pb.NewAccountServiceClient(c.accountConn).GetAccount(ctx, req)
}
