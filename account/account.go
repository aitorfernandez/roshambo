package main

import (
	"context"

	pb "github.com/aitorfernandez/roshambo/proto"
)

// accountServer is the API for AccountService.
type accountServer struct{}

func newAccountServer() *accountServer {
	return &accountServer{}
}

// GetAccount calls store for the given id.
func (s accountServer) GetAccount(ctx context.Context, req *pb.GetAccountReq) (*pb.Account, error) {
	return &pb.Account{
		ID:    req.ID,
		Email: "bar@gmail.com",
	}, nil
}
