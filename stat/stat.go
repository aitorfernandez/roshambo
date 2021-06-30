package main

import (
	"context"
	"math/rand"
	"time"

	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/aitorfernandez/roshambo/stat/model"
	"github.com/aitorfernandez/roshambo/stat/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// statServer is the API for statServer service.
type statServer struct {
	store *store.Store
}

func newStatServer(s *store.Store) *statServer {
	return &statServer{s}
}

// CreateStat creates a new stat wiht the given stat.
func (s statServer) CreateStat(ctx context.Context, req *pb.CreateStatReq) (*pb.Stat, error) {
	rand.Seed(time.Now().UnixNano())

	var (
		st = &model.Stat{
			AccountID:    req.Stat.AccountID,
			ComputerMove: int32(rand.Intn(2 + 1)),
			PlayerMove:   req.Stat.PlayerMove,
		}
		err error
	)
	if st, err = s.store.CreateStat(ctx, st); err != nil {
		return nil, err
	}
	return st.Proto(), nil
}

// DeleteStatsByAccount deletes the stats for the given account.
func (s statServer) DeleteStatsByAccount(ctx context.Context, req *pb.DeleteStatsByAccountReq) (*pb.DeleteStatsByAccountRes, error) {
	var (
		rows int32
		err  error
	)
	if rows, err = s.store.DeleteStatsByAccount(ctx, req.AccountID); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.DeleteStatsByAccountRes{
		RowsAffected: rows,
	}, nil
}

// ListStatsByAccount lists all stats for the given account.
func (s statServer) ListStatsByAccount(ctx context.Context, req *pb.ListStatsByAccountReq) (*pb.ListStatsByAccountRes, error) {
	var (
		stats []*model.Stat
		err   error
	)
	if stats, err = s.store.StatsByAccount(ctx, req.AccountID); err != nil {
		return nil, err
	}
	ss := make([]*pb.Stat, 0, len(stats))
	for _, s := range stats {
		ss = append(ss, s.Proto())
	}
	return &pb.ListStatsByAccountRes{
		Stats: ss,
	}, nil
}

// ListRankings lists all rankings.
func (s statServer) ListRankings(ctx context.Context, req *emptypb.Empty) (*pb.ListRankingsRes, error) {
	var (
		rankings []*model.Ranking
		err      error
	)
	if rankings, err = s.store.Rankings(ctx); err != nil {
		return nil, err
	}
	rr := make([]*pb.Ranking, 0, len(rankings))
	for _, r := range rankings {
		rr = append(rr, r.Proto())
	}
	return &pb.ListRankingsRes{
		Rankings: rr,
	}, nil
}
