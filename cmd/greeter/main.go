package main

import (
	"fmt"
	"log"
	"net"

	"github.com/n7down/microservices/internal/greeter/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = "8081"
	cert = "cert.pem"
	key  = "key.pem"
)

type greeterServer struct{}

func (g *greeterServer) SayHello(ctx context.Context, req *greeter_pb.HelloRequest) (*greeter_pb.HelloResponse, error) {
	return &greeter_pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	fmt.Printf("Listening on port: %s\n", port)
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	greeter_pb.RegisterGreeterServiceServer(grpcServer, &greeterServer{})
	grpcServer.Serve(lis)
}
