package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/n7down/microservices/internal/greeter/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	port = "8081"
)

type greeterServer struct{}

func (g *greeterServer) SayHello(ctx context.Context, req *greeter_pb.HelloRequest) (*greeter_pb.HelloResponse, error) {
	return &greeter_pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Listening on port: %s\n", port)
	grpcServer := grpc.NewServer()
	greeter_pb.RegisterHelloServiceServer(grpcServer, &greeterServer{})
	grpcServer.Serve(lis)
}
