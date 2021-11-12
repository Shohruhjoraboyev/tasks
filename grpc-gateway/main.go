package main

import (
	"context"
	"fmt"
	"gateway/proto/service"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, req *service.Req) (resp *service.Resp, err error) {
	return &service.Resp{Message: req.Name + " world"}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	service.RegisterGreeterServer(s, &service.UnimplementedGreeterServer{})

	log.Println("Serving gRPC on 0.0.0.0:8081")
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8081",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	err = service.RegisterGreeterHandler(context.Background(), gwmux, conn)

	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	ser, err := NewServer().SayHello(context.Background(), &service.Req{Name: "hello"})
	if err != nil {
		log.Fatalln("Failed implementing SayHello:", err)
	}
	fmt.Println(ser)
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())

}
