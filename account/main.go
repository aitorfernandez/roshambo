package main

import (
	"log"
	"net"

	"github.com/aitorfernandez/roshambo/account/store"
	"github.com/aitorfernandez/roshambo/pkg/env"
	pb "github.com/aitorfernandez/roshambo/proto"
	"google.golang.org/grpc"
)

func die(err error) {
	log.Fatalf("account service %s", err.Error())
}

func main() {
	var (
		err error
		lis net.Listener
		sto *store.Store
	)
	if lis, err = net.Listen("tcp", env.MustHget("account", "addr")); err != nil {
		die(err)
	}
	if sto, err = store.New(env.MustHget("account", "psql")); err != nil {
		die(err)
	}

	srv := grpc.NewServer()
	pb.RegisterAccountServiceServer(srv, newAccountServer(sto))
	if err = srv.Serve(lis); err != nil {
		die(err)
	}
}
