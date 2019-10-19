package users

import (
	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n7down/microservices/internal/client/users/pb"
	"github.com/n7down/microservices/internal/client/users/request"
	"github.com/n7down/microservices/internal/client/users/response"
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

	r, err := s.usersClient.UsersGetPassword(c, &users_pb.UsersGetPasswordRequest{Username: req.Username})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if r == nil {
		c.JSON(http.StatusNoContent, res)
		return
	}

	res = response.CheckPasswordResponse{
		ID:       r.ID,
		Password: r.Password,
	}

	c.JSON(http.StatusOK, res)
}

func (s *UsersServer) CreateHandler(c *gin.Context) {
	var (
		req request.CreateRequest
		res response.CreateResponse
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

	r, err := s.usersClient.UsersCreate(c, &users_pb.UsersCreateRequest{
		Username:  req.Username,
		Password:  req.Password,
		Firstname: req.Firstname,
		Lastname:  req.Lastname})

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = response.CreateResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}

func (s *UsersServer) ByIdHandler(c *gin.Context) {
}

// FIXME: update this
func (s *UsersServer) ListHandler(c *gin.Context) {
	var (
		res response.ListResponse
	)

	r, err := s.usersClient.UsersList(c, &users_pb.UsersListRequest{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if len(r.Users) > 0 {
		c.JSON(http.StatusNoContent, res)
		return
	}

	// FIXME: update this
	res = response.ListResponse{
		//Message: r.Message,
	}

	c.JSON(http.StatusOK, res)
}

// FIXME: implement
func (s *UsersServer) UpdateHandler(c *gin.Context) {
}

// FIXME: implement
func (s *UsersServer) DeleteHandler(c *gin.Context) {
	//id := c.Query("id")
}
