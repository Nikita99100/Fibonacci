package server

import (
	"context"
	"github.com/Nikita99100/Fibonacci/handler"
	//"github.com/Nikita99100/Fibonacci/handler"
	"github.com/Nikita99100/Fibonacci/pkg/api"
)

type GRPCServer struct {
	Handler *handler.Handler
	api.UnimplementedFibonacciServer
}

func (s *GRPCServer) FibonacciSequence(ctx context.Context, req *api.FibonacciRequest) (*api.FibonacciResponse, error) {
	result, err := s.Handler.GetFibonacciSequence(int(req.GetX()), int(req.GetY()))
	if err != nil {
		return nil, err
	}
	return &api.FibonacciResponse{Result: result}, nil
}
