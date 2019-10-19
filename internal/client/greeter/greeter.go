package greeter

import (
	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n7down/microservices/internal/client/greeter/request"
	"github.com/n7down/microservices/internal/client/greeter/response"
	"github.com/n7down/microservices/internal/pb/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	cert = "cert.pem"
)

type GreeterServer struct {
	greeterClient greeter_pb.GreeterServiceClient
}

func NewGreeterServer(serverEnv string) (*GreeterServer, error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	greeterConn, err := grpc.Dial(serverEnv, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		return nil, err
	}
	server := &GreeterServer{
		greeterClient: greeter_pb.NewGreeterServiceClient(greeterConn),
	}
	return server, nil
}

func (s *GreeterServer) HelloHandler(c *gin.Context) {

	// FIXME: need to read up on this - https://blog.golang.org/context
	//clientCtx, cancel := context.WithTimeout(c, time.Second)
	//defer cancel()

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

	r, err := s.greeterClient.SayHello(c, &greeter_pb.HelloRequest{Name: req.Name})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = response.HelloResponse{
		Message: r.Message,
	}

	c.JSON(http.StatusOK, res)
}
