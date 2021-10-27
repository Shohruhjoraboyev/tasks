package main

import (
	"GOPATH/grpc/pkg/contacts"
	"GOPATH/grpc/pkg/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &service.GRPCServer{}
	contacts.RegisterContactServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
