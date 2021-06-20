package main

import (
	"log"
	"net"

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
	)
	if lis, err = net.Listen("tcp", ":5050"); err != nil {
		die(err)
	}

	srv := grpc.NewServer()
	pb.RegisterAccountServiceServer(srv, newAccountServer())
	if err = srv.Serve(lis); err != nil {
		die(err)
	}
}
