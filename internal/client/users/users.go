package users

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/n7down/microservices/internal/client/users/request"
	"github.com/n7down/microservices/internal/client/users/response"
	"github.com/n7down/microservices/internal/pb/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	cert        = "cert.pem"
	FIVEMINUTES = 5 * time.Minute
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
	ctx, cancel := context.WithTimeout(c, FIVEMINUTES)
	defer cancel()

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

	r, err := s.usersClient.UsersGetPassword(ctx, &users_pb.UsersGetPasswordRequest{Username: req.Username})
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
		Lastname:  req.Lastname,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = response.CreateResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}

func (s *UsersServer) ByIDHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVEMINUTES)
	defer cancel()

	var (
		id  string = c.Param("id")
		req request.ByIDRequest
		res response.ByIDResponse
	)

	req.ID = id

	// validate input data
	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := s.usersClient.UsersByID(ctx, &users_pb.UsersByIDRequest{
		ID: id,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = response.ByIDResponse{
		ID:        r.ID,
		Username:  r.Username,
		Firstname: r.Firstname,
		Lastname:  r.Lastname,
	}

	c.JSON(http.StatusOK, res)
}

func (s *UsersServer) ListHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVEMINUTES)
	defer cancel()

	var (
		res response.ListResponse
	)

	stream, err := s.usersClient.UsersList(ctx, &users_pb.UsersListRequest{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		user := response.User{
			ID:        r.ID,
			Username:  r.Username,
			Firstname: r.Firstname,
			Lastname:  r.Lastname,
		}
		res.Users = append(res.Users, user)
	}

	if len(res.Users) == 0 {
		c.JSON(http.StatusNoContent, res)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (s *UsersServer) UpdateHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVEMINUTES)
	defer cancel()

	var (
		id  string = c.Param("id")
		req request.UpdateRequest
		res response.UpdateResponse
	)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	req.ID = id

	// validate input data
	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := s.usersClient.UsersUpdate(ctx, &users_pb.UsersUpdateRequest{
		ID:        req.ID,
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = response.UpdateResponse{
		ID:        r.ID,
		Username:  r.Username,
		Firstname: r.Firstname,
		Lastname:  r.Lastname,
	}

	c.JSON(http.StatusOK, res)
}

func (s *UsersServer) DeleteHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, FIVEMINUTES)
	defer cancel()

	var (
		id  string = c.Param("id")
		res response.DeleteResponse
	)

	req := request.DeleteRequest{
		ID: id,
	}

	// validate input data
	if validationErrors := req.Validate(); len(validationErrors) > 0 {
		err := map[string]interface{}{"validationError": validationErrors}
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r, err := s.usersClient.UsersDelete(ctx, &users_pb.UsersDeleteRequest{
		ID: req.ID,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = response.DeleteResponse{
		ID: r.ID,
	}

	c.JSON(http.StatusOK, res)
}
