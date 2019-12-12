package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/n7down/microservices/internal/apigateway"
	"github.com/n7down/microservices/internal/client/greeter"
	"github.com/n7down/microservices/internal/client/users"
	"github.com/n7down/microservices/internal/utils"
	log "github.com/sirupsen/logrus"
)

var (
	gatewayPort = "8080"
)

func main() {
	router := gin.Default()
	log.SetReportCaller(true)

	greeterServer, err := greeter.NewGreeterClient(utils.GetEnv("GREETER_HOST", "127.0.0.1:8081"))
	if err != nil {
		log.Fatal(err)
		return
	}

	usersServer, err := users.NewUsersClient(utils.GetEnv("USERS_HOST", "127.0.0.1:8082"))
	if err != nil {
		log.Fatal(err)
		return
	}

	apiGateway := apigateway.NewAPIGateway(
		greeterServer,
		usersServer,
	)

	authMiddleware, err := apiGateway.InitAuthRoutes(router)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = apiGateway.InitV1Routes(router, authMiddleware)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = apiGateway.InitV2Routes(router, authMiddleware)
	if err != nil {
		log.Fatal(err)
		return
	}

	routerPort := fmt.Sprintf(":%s", gatewayPort)
	log.Infof("Listening on port: %s\n", gatewayPort)
	err = apiGateway.Run(router, routerPort)
	if err != nil {
		log.Fatal(err)
		return
	}
}
