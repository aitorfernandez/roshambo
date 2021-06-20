package client

import (
	"context"
	"time"

	"github.com/aitorfernandez/roshambo/pkg/env"
	"google.golang.org/grpc"
)

// Client wrapps a collection of gRPC clients.
type Client struct {
	accountConn *grpc.ClientConn
}

// New creates a new Client.
func New() (*Client, error) {
	var (
		c   = &Client{}
		err error
	)
	if err = connGRPC(&c.accountConn, env.MustHget("account", "addr")); err != nil {
		return nil, err
	}
	return c, nil
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
