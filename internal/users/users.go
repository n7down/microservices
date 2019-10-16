package users

import (
	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n7down/microservices/internal/users/pb"
	"github.com/n7down/microservices/internal/users/request"
	"github.com/n7down/microservices/internal/users/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	cert = "cert.pem"
)

type UsersServer struct {
	usersClient users_pb.UsersServiceClient
}

func NewUsersServer(serverEnv string) (*UsersServer, error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	userConn, err := grpc.Dial(serverEnv, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		return nil, err
	}
	server := &UsersServer{
		usersClient: users_pb.NewUsersServiceClient(userConn),
	}
	return server, nil
}

// TODO: validation step for each handler
func (s *UsersServer) CheckPassword(c *gin.Context) {
	var (
		req request.CheckPasswordRequest
		res response.CheckPasswordResponse
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

	//r, err := s.greeterClient.SayHello(c, &greeter_pb.HelloRequest{Name: req.Name})
	//if err != nil {
	//c.JSON(http.StatusBadRequest, err)
	//return
	//}

	res = response.CheckPasswordResponse{
		//Message: r.Message,
	}

	c.JSON(http.StatusOK, res)
}

func (s *UsersServer) CreateHandler(c *gin.Context) {
}

func (s *UsersServer) ByIdHandler(c *gin.Context) {
}

func (s *UsersServer) ListHandler(c *gin.Context) {
	var (
		req request.ListRequest
		res response.ListResponse
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

	//r, err := s.greeterClient.SayHello(c, &greeter_pb.HelloRequest{Name: req.Name})
	//if err != nil {
	//c.JSON(http.StatusBadRequest, err)
	//return
	//}

	// FIXME: return no content
	//if len(r) > 0 {
	//c.JSON(http.StatusNoContent, res)
	//return
	//}

	res = response.ListResponse{
		//Message: r.Message,
	}

	c.JSON(http.StatusOK, res)
}

func (s *UsersServer) UpdateHandler(c *gin.Context) {
}

func (s *UsersServer) DeleteHandler(c *gin.Context) {
}
