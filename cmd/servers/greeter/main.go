package main

import (
	"fmt"
	"net"

	"github.com/n7down/microservices/internal/pb/greeter"
	server "github.com/n7down/microservices/internal/servers/greeter"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = "8081"
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

	log.Infof("Listening on port: %s\n", port)
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	greeter_pb.RegisterGreeterServiceServer(grpcServer, &server.GreeterServer{})
	grpcServer.Serve(lis)
}
