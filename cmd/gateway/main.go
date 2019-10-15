package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/n7down/microservices/internal/gateway"
	"github.com/n7down/microservices/internal/greeter"
	"github.com/n7down/microservices/internal/users"
	log "github.com/sirupsen/logrus"
)

var (
	gatewayPort = "8080"
)

func main() {
	router := gin.New()
	log.SetReportCaller(true)

	greeterServer, err := greeter.NewGreeterServer(os.Getenv("GREETER_HOST"))
	if err != nil {
		log.Fatal(err)
		return
	}

	usersServer, err := users.NewUsersServer(os.Getenv("USERS_HOST"))
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
