package main

import (
	"fmt"
	"net"

	"github.com/n7down/microservices/internal/pb/users"
	servers "github.com/n7down/microservices/internal/servers/users"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = "8082"
	cert = "cert.pem"
	key  = "key.pem"
)

func main() {
	log.SetReportCaller(true)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	usersServer, err := servers.NewUsersServer()
	if err != nil {
		log.Fatal("unable to start users server")
	}

	log.Info("Listening on port: %s\n", port)
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	users_pb.RegisterUsersServiceServer(grpcServer, usersServer)
	grpcServer.Serve(lis)
}
