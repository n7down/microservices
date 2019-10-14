package greeter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n7down/microservices/internal/greeter/pb"
	"github.com/n7down/microservices/internal/greeter/request"
	"github.com/n7down/microservices/internal/greeter/response"
)

type Greeter struct {
	Client greeter_pb.GreeterServiceClient
}

func NewGreeter(c greeter_pb.GreeterServiceClient) *Greeter {
	return &Greeter{Client: c}
}

func (g Greeter) HelloHandler(c *gin.Context) {
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

	r, err := g.Client.SayHello(c, &greeter_pb.HelloRequest{Name: req.Name})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res = response.HelloResponse{
		Message: r.Message,
	}

	c.JSON(http.StatusOK, res)
}
