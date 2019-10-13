package messages

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/n7down/microservices/internal/greeter/pb"
)

const (
	messagesPort = "8081"
)

type Messages struct {
	Client greeter_pb.HelloServiceClient
}

func NewMessages(c greeter_pb.HelloServiceClient) *Messages {
	return &Messages{Client: c}
}

type HelloMessageRequest struct {
	Name string `json: "name" binding:"required"`
}

func (r *HelloMessageRequest) Validate() url.Values {
	errs := url.Values{}
	if r.Name == "" {
		errs.Add("name", "The name field is required!")
	}
	return errs
}

type HelloMessageResponse struct {
	Message string `json: "message"`
}

func (m Messages) HelloHandler(c *gin.Context) {
	var (
		req HelloMessageRequest
		res HelloMessageResponse
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

	r, err := m.Client.SayHello(c, &greeter_pb.HelloRequest{Name: req.Name})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = HelloMessageResponse{
		Message: r.Message,
	}

	c.JSON(http.StatusOK, res)
}
