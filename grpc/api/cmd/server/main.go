package main

import (
	"GOPATH/grpc/pkg/contacts"
	"GOPATH/grpc/pkg/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	srv := &service.GRPCServer{}
	s := grpc.NewServer()

	contacts.RegisterContactServiceServer(s, srv)

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
