package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aitorfernandez/roshambo/profile/store"
	pb "github.com/aitorfernandez/roshambo/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// profileServer is the API for profileServer service.
type profileServer struct {
	store *store.Store
}

func newProfileServer(s *store.Store) *profileServer {
	return &profileServer{s}
}

// GetProfileByAccount gets a profile for the given accountID.
func (s profileServer) GetProfileByAccount(ctx context.Context, req *pb.GetProfileByAccountReq) (*pb.Profile, error) {
	p, err := s.store.ProfileByAccount(ctx, req.AccountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return p.Proto(), nil
}
