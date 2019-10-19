package main

import (
	"fmt"
	"log"
	"net"

	servers "github.com/n7down/microservices/internal/servers/users"
	"github.com/n7down/microservices/internal/users/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = "8082"
	cert = "cert.pem"
	key  = "key.pem"
)

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
	users_pb.RegisterUsersServiceServer(grpcServer, &servers.UsersServer{})
	grpcServer.Serve(lis)
}
