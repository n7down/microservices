package messages

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/n7down/microservices/internal/messages/pb"
)

const (
	messagesPort = "8081"
)

type Messages struct {
	Client messages.HelloServiceClient
}

func NewMessages(c messages.HelloServiceClient) *Messages {
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

	// FIXME: change package messages for HelloRequest to messages.pb
	r, err := m.Client.SayHello(c, &messages.HelloRequest{Name: req.Name})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = HelloMessageResponse{
		Message: r.Message,
	}

	c.JSON(http.StatusOK, res)
}
