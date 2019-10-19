package server

import (
	"context"

	"github.com/n7down/microservices/internal/pb/greeter"
)

type GreeterServer struct{}

func (g *GreeterServer) SayHello(ctx context.Context, req *greeter_pb.HelloRequest) (*greeter_pb.HelloResponse, error) {
	return &greeter_pb.HelloResponse{Message: "Hello " + req.Name}, nil
}
