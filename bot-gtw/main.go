package main

import (
	"bot-gtw/handler"
	"bot-gtw/proto/sender"
	"log"

	"google.golang.org/grpc"
)

func main() {
	srv, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := sender.NewSendClient(srv)
	handler.Init(c)
}
