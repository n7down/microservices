package gateway

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"github.com/n7down/microservices/internal/gateway/request"
	"github.com/n7down/microservices/internal/gateway/response"

	"github.com/n7down/microservices/internal/greeter"
	"github.com/n7down/microservices/internal/greeter/pb"

	"github.com/n7down/microservices/internal/users"
)

type Gateway struct {
	greeterServiceClient greeter_pb.GreeterServiceClient
}

func NewGateway(greeterServiceClient greeter_pb.GreeterServiceClient) *Gateway {
	return &Gateway{
		greeterServiceClient: greeterServiceClient,
	}
}

func (g Gateway) authLogin(in request.LoginRequest) (*response.LoginResponse, error) {
	//var user users.User

	// FIXME: get the user
	//err := user.GetUserByID(in.Username)
	//if err != nil {
	//return &auth.LoginResponse{}, err
	//}

	//if isValidSecret, _ := utils.CheckUserSecret(user.HashSecret, in.Password); isValidSecret {
	//return &auth.LoginResponse{
	//FirstName:       user.FirstName,
	//LastName:        user.LastName,
	//UserID:          user.UserID,
	//Secret:          user.Secret,
	//IDType:          user.IDType,
	//SecretValidFrom: user.SecretValidFrom,
	//SecretValidTo:   user.SecretValidTo,
	//}, nil
	//}

	if in.Username == "leslie2020" && in.Password == "password" {
		return &response.LoginResponse{
			FirstName: "Leslie",
			LastName:  "Knope",
			Username:  "leslie2020",
		}, nil
	}

	return &response.LoginResponse{}, jwt.ErrFailedAuthentication
}

func (g Gateway) InitAuthRoutes(r *gin.Engine) (*jwt.GinJWTMiddleware, error) {
	var loginResponse *response.LoginResponse

	athMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "production",
		Key:              []byte(os.Getenv("AUTH_KEY")),
		Timeout:          time.Duration(24) * time.Hour,
		MaxRefresh:       time.Duration(24) * time.Hour,
		IdentityKey:      "id",
		SigningAlgorithm: "HS256",
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var in request.LoginRequest
			if err := c.ShouldBind(&in); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			response := &response.LoginResponse{}
			response, err := g.authLogin(in)
			loginResponse = response

			if err != nil {
				return "", err
			}

			return response, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		LoginResponse: func(c *gin.Context, statusCode int, token string, tokenExpires time.Time) {
			loginResponse.Token = token
			loginResponse.Expires = tokenExpires.String()
			c.JSON(http.StatusOK, loginResponse)
		},

		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	})
	return authMiddleware, err
}

func (g Gateway) InitRoutes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) error {
	v1 := r.Group("/api/v1")

	auth := v1.Group("/auth")
	auth.GET("/refresh", authMiddleware.RefreshHandler)
	auth.POST("/login", authMiddleware.LoginHandler)

	greeterGroup := v1.Group("/greeter")
	{
		g := greeter.NewGreeter(g.greeterServiceClient)
		greeterGroup.GET("/hello", g.HelloHandler)
	}

	u := users.NewUsers()
	v1.POST("/users/create", u.CreateHandler)

	usersGroup := v1.Group("/users")
	usersGroup.Use(authMiddleware.MiddlewareFunc())
	{
		usersGroup.GET("/list", u.ListHandler)
		usersGroup.GET("/byid/:id", u.ByIdHandler)
		usersGroup.PUT("/update/:id", u.UpdateHandler)
		usersGroup.DELETE("/delete/:id", u.DeleteHandler)
	}

	products := v1.Group("/products")
	products.Use(authMiddleware.MiddlewareFunc())
	{
	}

	baskets := v1.Group("/baskets")
	baskets.Use(authMiddleware.MiddlewareFunc())
	{
	}

	activity := v1.Group("/activity")
	activity.Use(authMiddleware.MiddlewareFunc())
	{
	}
	return nil
}

func (g Gateway) Run(router *gin.Engine, routerPort string) error {
	err := http.ListenAndServe(routerPort, router)
	if err != nil {
		return err
	}
	return nil
}
