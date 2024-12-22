package main

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/aitorfernandez/roshambo/account/model"
	"github.com/aitorfernandez/roshambo/account/store"
	"github.com/aitorfernandez/roshambo/pkg/env"
	"github.com/aitorfernandez/roshambo/pkg/uid"
	pb "github.com/aitorfernandez/roshambo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// accountServer is the API for AccountService.
type accountServer struct {
	mailConn *grpc.ClientConn
	store    *store.Store
}

func newAccountServer(s *store.Store) (*accountServer, error) {
	var (
		a = &accountServer{
			store: s,
		}
		err error
	)
	if err = connGRPC(&a.mailConn, env.MustHget("mail", "addr")); err != nil {
		return nil, err
	}
	return a, nil
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
		err   error
		token string
	)
	if err = a.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if a, err = s.store.SetAccount(ctx, a); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if token, err = a.GenerateToken(time.Now().Unix(), req.IP); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	go func() {
		pb.NewMailServiceClient(s.mailConn).SendToken(context.Background(), &pb.SendTokenReq{
			ID:       a.ID,
			Receiver: a.Email,
			Token:    token,
		})
	}()

	return &wrapperspb.BoolValue{
		Value: true,
	}, nil
}

func connGRPC(conn **grpc.ClientConn, addr string) error {
	var (
		err  error
		opts []grpc.DialOption
	)
	opts = append(
		opts,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second*5),
	)
	if *conn, err = grpc.DialContext(context.Background(), addr, opts...); err != nil {
		return err
	}
	return nil
}

// ValidateToken validates GetStarted token.
func (s accountServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenReq) (*pb.Account, error) {
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
	if ok := a.ValidateToken(req.Token, req.IP); !ok {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	// UpdateLastRequestAt invalid the token for use again.
	if a, err = s.store.UpdateLastRequestAt(ctx, a.ID, time.Now().Unix()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return a.Proto(), nil
}
