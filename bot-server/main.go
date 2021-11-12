package main

import (
	"bot-gtwy/handlers"
	"bot-gtwy/proto/sender"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	s := grpc.NewServer()
	srv := &handlers.Server{}

	sender.RegisterSendServer(s, srv)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	log.Println("Serving grpc on 0.0.0.0:8080")
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
