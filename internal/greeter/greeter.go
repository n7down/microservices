package greeter

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/n7down/microservices/internal/greeter/pb"
	"github.com/n7down/microservices/internal/greeter/request"
	"github.com/n7down/microservices/internal/greeter/response"
	"google.golang.org/grpc"
)

type GreeterServer struct {
	greeterClient greeter_pb.GreeterServiceClient
}

func NewGreeterServer(serverEnv string) (*GreeterServer, error) {
	greeterConn, err := grpc.Dial(serverEnv, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	server := &GreeterServer{
		greeterClient: greeter_pb.NewGreeterServiceClient(greeterConn),
	}
	return server, nil
}

func (s *GreeterServer) HelloHandler(c *gin.Context) {
	clientCtx, cancel := context.WithTimeout(c, time.Second)
	defer cancel()

	var (
		req request.HelloRequest
		res response.HelloResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// validate input data
	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := s.greeterClient.SayHello(clientCtx, &greeter_pb.HelloRequest{Name: req.Name})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = response.HelloResponse{
		Message: r.Message,
	}

	c.JSON(http.StatusOK, res)
}
