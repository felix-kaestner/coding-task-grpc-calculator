package server

import (
	"context"

	"github.com/felix-kaestner/calculator/internal/pkg/math"
	"github.com/felix-kaestner/calculator/internal/pkg/proto"
	"google.golang.org/grpc"
)

// server is used to implement the CalculatorServer interface
type server struct {
	proto.UnimplementedCalculatorServer
}

// Solve implements CalculatorServer interface
func (*server) Solve(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	res, err := math.Solve(req.GetMethod(), req.GetOperandA(), req.GetOperandB())
	if err != nil {
		return nil, err
	}

	return &proto.Response{Result: res}, nil
}

// New returns a new grpc server with the caluclator service registered
func New() *grpc.Server {
	s := grpc.NewServer()
	proto.RegisterCalculatorServer(s, &server{})
	return s
}
