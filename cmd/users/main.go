package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8081, "The server port")
)

func main() {
	flag.Parse()
	listenerPort := fmt.Sprintf(":%d", *port)
	lis, err := net.Listen("tcp", listenerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpcServer.Serve(lis)
	fmt.Sprintf("Listening on port %s", listenerPort)
}
