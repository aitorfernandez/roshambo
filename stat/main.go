package main

import (
	"log"
	"net"

	"github.com/aitorfernandez/roshambo/pkg/env"
	pb "github.com/aitorfernandez/roshambo/proto"
	"github.com/aitorfernandez/roshambo/stat/store"
	"google.golang.org/grpc"
)

func die(err error) {
	log.Fatalf("stat service %s", err.Error())
}

func main() {
	var (
		err error
		lis net.Listener
		s   *store.Store
	)
	if lis, err = net.Listen("tcp", env.MustHget("stat", "addr")); err != nil {
		die(err)
	}
	if s, err = store.New(env.MustHget("stat", "psql")); err != nil {
		die(err)
	}

	srv := grpc.NewServer()
	pb.RegisterStatServiceServer(srv, newStatServer(s))
	if err = srv.Serve(lis); err != nil {
		die(err)
	}
}
