package main

import (
	"log"
	"net"

	"github.com/aitorfernandez/roshambo/pkg/env"
	"github.com/aitorfernandez/roshambo/profile/store"
	pb "github.com/aitorfernandez/roshambo/proto"
	"google.golang.org/grpc"
)

func die(err error) {
	log.Fatalf("profile service %s", err.Error())
}

func main() {
	var (
		err error
		lis net.Listener
		s   *store.Store
	)
	if lis, err = net.Listen("tcp", env.MustHget("profile", "addr")); err != nil {
		die(err)
	}
	if s, err = store.New(env.MustHget("profile", "psql")); err != nil {
		die(err)
	}

	srv := grpc.NewServer()
	pb.RegisterProfileServiceServer(srv, newProfileServer(s))
	if err = srv.Serve(lis); err != nil {
		die(err)
	}
}
