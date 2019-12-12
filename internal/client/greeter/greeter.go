package greeter

import (
	"context"
	//"crypto/tls"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/n7down/microservices/internal/client/greeter/request"
	"github.com/n7down/microservices/internal/client/greeter/response"
	"github.com/n7down/microservices/internal/pb/greeter"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
)

const (
	cert        = "cert.pem"
	FIVEMINUTES = 5 * time.Minute
)

type GreeterClient struct {
	greeterClient greeter_pb.GreeterServiceClient
}

func NewGreeterClient(serverEnv string) (*GreeterClient, error) {
	//config := &tls.Config{
	//InsecureSkipVerify: true,
	//}

	//greeterConn, err := grpc.Dial(serverEnv, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	greeterConn, err := grpc.Dial(serverEnv, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	server := &GreeterClient{
		greeterClient: greeter_pb.NewGreeterServiceClient(greeterConn),
	}
	return server, nil
}

func (s *GreeterClient) HelloHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVEMINUTES)
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

	r, err := s.greeterClient.SayHello(ctx, &greeter_pb.HelloRequest{Name: req.Name})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = response.HelloResponse{
		Message: r.Message,
	}

	c.JSON(http.StatusOK, res)
}
