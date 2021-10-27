package adder

import (
	"GOPATH/grpcAdder/pkg/api"
	"context"
)

// GRPCServer ...

type GRPCServer struct{}

// Add ...

func (s *GRPCServer) Add(ctx context.Context, req *api.AddReq) (*api.AddResp, error) {
	return &api.AddResp{Result: req.GetX() + req.GetY()}, nil
}
