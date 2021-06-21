package main

import (
	"log"
	"net"

	"github.com/aitorfernandez/roshambo/pkg/env"
	pb "github.com/aitorfernandez/roshambo/proto"
	"google.golang.org/grpc"
)

func die(err error) {
	log.Fatalf("mail service %s", err.Error())
}

// main starts the gRPC mail server.
func main() {
	var (
		err error
		lis net.Listener
	)
	if lis, err = net.Listen("tcp", env.MustHget("mail", "addr")); err != nil {
		die(err)
	}

	srv := grpc.NewServer()
	pb.RegisterMailServiceServer(srv, &mailServer{})
	if err = srv.Serve(lis); err != nil {
		die(err)
	}
}
