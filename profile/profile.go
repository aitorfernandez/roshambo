package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aitorfernandez/roshambo/pkg/uid"
	"github.com/aitorfernandez/roshambo/profile/model"
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

// SetProfile updates the given profile.
func (s profileServer) SetProfile(ctx context.Context, req *pb.SetProfileReq) (*pb.Profile, error) {
	var (
		p = &model.Profile{
			ID:        uid.New(model.ProfilePrefix),
			AccountID: req.Profile.AccountID,
			Avatar:    req.Profile.Avatar,
			Username:  req.Profile.Username,
		}
		err error
	)
	if err = p.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	p, err = s.store.SetProfile(ctx, p)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return p.Proto(), nil
}
