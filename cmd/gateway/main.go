package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/n7down/microservices/internal/gateway"

	"github.com/n7down/microservices/internal/greeter/pb"
)

var (
	gatewayPort = "8080"
)

func main() {
	router := gin.New()
	log.SetReportCaller(true)

	// greeter
	greeterConn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer greeterConn.Close()
	greeterClient := greeter_pb.NewGreeterServiceClient(greeterConn)

	gateway := gateway.NewGateway(
		greeterClient,
	)

	authMiddleware, err := gateway.InitAuthRoutes(router)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = gateway.InitRoutes(router, authMiddleware)
	if err != nil {
		log.Fatal(err)
		return
	}

	routerPort := fmt.Sprintf(":%s", gatewayPort)
	fmt.Printf("Listening on port: %s\n", gatewayPort)
	err = gateway.Run(router, routerPort)
	if err != nil {
		log.Fatal(err)
		return
	}
}
