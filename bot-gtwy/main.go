package main

import (
	"bot-gtwy/handlers"
	"bot-gtwy/proto/sender"
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

type GRPCServer struct{}

func NewServer() *GRPCServer {
	return &GRPCServer{}
}
func (s *GRPCServer) SendMsg(ctx context.Context, req *sender.SendReq) (*sender.SendResp, error) {
	err := handlers.SendMessage()
	return &sender.SendResp{Status: "message send"}, err
}

var resp string

func (s *GRPCServer) GetMsgs(ctx context.Context, req *sender.GetReq) (*sender.GetResp, error) {
	msgsarr := handlers.SortByPriority()
	for _, msgs := range msgsarr {
		for _, msg := range msgs {
			resp = resp + " priority: " + msg.Priority + " text: " + msg.Text + ";"
		}
	}

	// for _, msgs := range msgsarr {
	// 	for _, msg := range msgs {
	// 		botMessage := messg{Text: msg.Text}
	// 		resp, err := json.Marshaler(botMessage)
	// 	}

	// }

	return &sender.GetResp{Text: resp}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	srv := &GRPCServer{}
	s := grpc.NewServer()
	sender.RegisterSendServer(s, srv)
	log.Println("Serving grpc on 0.0.0.0:8080")
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = sender.RegisterSendHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	log.Println("Serving gRPC-Gsteway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
