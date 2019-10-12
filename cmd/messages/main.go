package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/n7down/microservices/internal/messages/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	port = "8081"
)

type testServer struct{}

func (t *testServer) SayHello(ctx context.Context, req *messages.HelloRequest) (*messages.HelloResponse, error) {
	return &messages.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Listening on port: %s\n", port)
	grpcServer := grpc.NewServer()
	messages.RegisterHelloServiceServer(grpcServer, &testServer{})
	grpcServer.Serve(lis)
}
