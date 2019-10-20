package main

import (
	"fmt"
	//"os"

	"github.com/gin-gonic/gin"
	"github.com/n7down/microservices/internal/client/greeter"
	"github.com/n7down/microservices/internal/client/users"
	"github.com/n7down/microservices/internal/gateway"
	"github.com/n7down/microservices/internal/utils"
	log "github.com/sirupsen/logrus"
)

var (
	gatewayPort = "8080"
)

func main() {
	router := gin.Default()
	log.SetReportCaller(true)

	//greeterServer, err := greeter.NewGreeterServer(os.Getenv("GREETER_HOST"))
	greeterServer, err := greeter.NewGreeterServer(utils.GetEnv("GREETER_HOST", "127.0.0.1:8081"))
	if err != nil {
		log.Fatal(err)
		return
	}

	//usersServer, err := users.NewUsersServer(os.Getenv("USERS_HOST"))
	usersServer, err := users.NewUsersServer(utils.GetEnv("USERS_HOST", "127.0.0.1:8082"))
	if err != nil {
		log.Fatal(err)
		return
	}

	gateway := gateway.NewGateway(
		greeterServer,
		usersServer,
	)

	authMiddleware, err := gateway.InitAuthRoutes(router)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = gateway.InitV1Routes(router, authMiddleware)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = gateway.InitV2Routes(router, authMiddleware)
	if err != nil {
		log.Fatal(err)
		return
	}

	routerPort := fmt.Sprintf(":%s", gatewayPort)
	log.Infof("Listening on port: %s\n", gatewayPort)
	err = gateway.Run(router, routerPort)
	if err != nil {
		log.Fatal(err)
		return
	}
}
