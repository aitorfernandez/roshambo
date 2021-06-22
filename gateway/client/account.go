package client

import (
	"context"

	pb "github.com/aitorfernandez/roshambo/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// GetAccount sends GetAccount request ro accountServer.
func (c Client) GetAccount(ctx context.Context, req *pb.GetAccountReq) (*pb.Account, error) {
	return pb.NewAccountServiceClient(c.accountConn).GetAccount(ctx, req)
}

// GetStarted sends GetStarted request to accountServer.
func (c Client) GetStarted(ctx context.Context, req *pb.GetStartedReq) (*wrapperspb.BoolValue, error) {
	return pb.NewAccountServiceClient(c.accountConn).GetStarted(ctx, req)
}

// ValidateToken sends ValidateAccess request to accountServer.
func (c Client) ValidateToken(ctx context.Context, req *pb.ValidateTokenReq) (*pb.Account, error) {
	return pb.NewAccountServiceClient(c.accountConn).ValidateToken(ctx, req)
}
