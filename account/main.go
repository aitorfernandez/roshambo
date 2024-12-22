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
		srvAcc *accountServer
		err    error
		lis    net.Listener
		s      *store.Store
	)
	if lis, err = net.Listen("tcp", env.MustHget("account", "addr")); err != nil {
		die(err)
	}
	if s, err = store.New(env.MustHget("account", "psql")); err != nil {
		die(err)
	}
	if srvAcc, err = newAccountServer(s); err != nil {
		die(err)
	}

	srv := grpc.NewServer()
	pb.RegisterAccountServiceServer(srv, srvAcc)
	if err = srv.Serve(lis); err != nil {
		die(err)
	}
}
