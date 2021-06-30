package client

import (
	"context"

	pb "github.com/aitorfernandez/roshambo/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateStat sends CreateStat request to statServer.
func (c Client) CreateStat(ctx context.Context, req *pb.CreateStatReq) (*pb.Stat, error) {
	return pb.NewStatServiceClient(c.statConn).CreateStat(ctx, req)
}

// DeleteStatsByAccount sends DeleteStatsByAccount request to statServer.
func (c Client) DeleteStatsByAccount(ctx context.Context, req *pb.DeleteStatsByAccountReq) (int32, error) {
	res, err := pb.NewStatServiceClient(c.statConn).DeleteStatsByAccount(ctx, req)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected, nil
}

// ListStatsByAccount sends ListStatsByAccount request to statServer.
func (c Client) ListStatsByAccount(ctx context.Context, req *pb.ListStatsByAccountReq) ([]*pb.Stat, error) {
	res, err := pb.NewStatServiceClient(c.statConn).ListStatsByAccount(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Stats, nil
}

// ListRankings sends ListRankings request to statServer.
func (c Client) ListRankings(ctx context.Context, req *emptypb.Empty) ([]*pb.Ranking, error) {
	res, err := pb.NewStatServiceClient(c.statConn).ListRankings(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Rankings, nil
}
