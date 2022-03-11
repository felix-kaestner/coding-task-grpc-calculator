package client

import (
	"context"
	"log"
	"strings"

	"github.com/felix-kaestner/calculator/internal/pkg/math"
	"github.com/felix-kaestner/calculator/internal/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// client is used to wrap the default implementation of the CalculatorClient
// interface and provide a more convenient API with defaults.
type client struct {
	cc proto.CalculatorClient
}

// Solve sends a mathematical operation to the server and returns the result by
// invoking the wrapped CalculatorClient implementation.
func (c *client) Solve(ctx context.Context, method string, a float64, b float64) (float64, error) {
	// validate the provided method
	op, ok := proto.Method_value[strings.ToUpper(method)]
	if !ok {
		return 0, math.ErrInvalidMethod
	}

	// Delegate to the wrapped CalculatorClient implementation.
	r, err := c.cc.Solve(ctx, &proto.Request{Method: proto.Method(op), OperandA: a, OperandB: b})
	if err != nil {
		return 0, err
	}

	// Unwrap the response.
	return r.Result, nil
}

// Solve sends a mathematical operation to the server and returns the result by
// internally creating a insecure grpc client connection to the specified host address.
func Solve(ctx context.Context, address string, method string, a float64, b float64) (float64, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	return New(conn).Solve(ctx, method, a, b)
}

// New returns a new grpc server with the caluclator service registered
func New(conn grpc.ClientConnInterface) *client {
	return &client{cc: proto.NewCalculatorClient(conn)}
}
