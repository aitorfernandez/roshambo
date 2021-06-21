package main

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/aitorfernandez/roshambo/account/model"
	"github.com/aitorfernandez/roshambo/account/store"
	"github.com/aitorfernandez/roshambo/pkg/uid"
	pb "github.com/aitorfernandez/roshambo/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// accountServer is the API for AccountService.
type accountServer struct {
	store *store.Store
}

func newAccountServer(s *store.Store) *accountServer {
	return &accountServer{
		store: s,
	}
}

// GetAccount calls store for the given id.
func (s accountServer) GetAccount(ctx context.Context, req *pb.GetAccountReq) (*pb.Account, error) {
	var (
		a   *model.Account
		err error
	)
	if a, err = s.store.Account(ctx, req.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return a.Proto(), nil
}

// GetStarted starts a sign in or sign up.
func (s accountServer) GetStarted(ctx context.Context, req *pb.GetStartedReq) (*wrapperspb.BoolValue, error) {
	var (
		a = &model.Account{
			ID:            uid.New(model.AccountPrefix),
			Email:         req.Email,
			LastRequestAt: time.Now().Unix(),
		}
		err error
	)
	if err = a.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if a, err = s.store.SetAccount(ctx, a); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &wrapperspb.BoolValue{
		Value: true,
	}, nil
}
