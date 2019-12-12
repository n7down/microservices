package apigateway

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"github.com/n7down/microservices/internal/apigateway/request"
	"github.com/n7down/microservices/internal/apigateway/response"

	"github.com/n7down/microservices/internal/client/greeter"
	"github.com/n7down/microservices/internal/client/users"
)

type APIGateway struct {
	greeterClient *greeter.GreeterClient
	usersClient   *users.UsersClient
}

func NewAPIGateway(g *greeter.GreeterClient, u *users.UsersClient) *APIGateway {
	return &APIGateway{
		greeterClient: g,
		usersClient:   u,
	}
}

func (g *APIGateway) authLogin(in request.LoginRequest) (*response.LoginResponse, error) {
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

	if in.Username == "admin" && in.Password == "password" {
		return &response.LoginResponse{
			FirstName: "Leslie",
			LastName:  "Knope",
			Username:  "admin",
		}, nil
	}

	return &response.LoginResponse{}, jwt.ErrFailedAuthentication
}

func (g *APIGateway) logResponseLatencyMiddleware(c *gin.Context) {
	start := time.Now()
	c.Next()

	// FIXME: log the end time
	_ = time.Now().Sub(start)
}

func (g *APIGateway) InitAuthRoutes(r *gin.Engine) (*jwt.GinJWTMiddleware, error) {
	const identityKey = "username"
	var loginResponse *response.LoginResponse

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "production",
		Key:              []byte(os.Getenv("AUTH_KEY")),
		Timeout:          time.Duration(24) * time.Hour,
		MaxRefresh:       time.Duration(24) * time.Hour,
		IdentityKey:      identityKey,
		SigningAlgorithm: "HS256",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*response.LoginResponse); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
					"firstname": v.FirstName,
					"lastname":  v.LastName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &response.LoginResponse{
				Username: claims[identityKey].(string),
			}
		},
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

func (g *APIGateway) InitV1Routes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) error {
	v1 := r.Group("/api/v1")
	v1.Use(g.logResponseLatencyMiddleware)

	auth := v1.Group("/auth")
	{
		auth.GET("/refresh", authMiddleware.RefreshHandler)
		auth.POST("/login", authMiddleware.LoginHandler)
	}

	greeterGroup := v1.Group("/greeter")
	{
		greeterGroup.GET("/hello", g.greeterClient.HelloHandler)
	}

	usersGroup := v1.Group("/users")
	usersGroup.POST("/create", g.usersClient.CreateHandler)
	//usersGroup.Use(authMiddleware.MiddlewareFunc())
	{
		usersGroup.GET("/list", g.usersClient.ListHandler)
		usersGroup.GET("/byid/:id", g.usersClient.ByIDHandler)
		usersGroup.PUT("/update/:id", g.usersClient.UpdateHandler)
		usersGroup.DELETE("/delete/:id", g.usersClient.DeleteHandler)
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

// FIXME: implement - use graphql
func (g *APIGateway) InitV2Routes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) error {
	return nil
}

func (g *APIGateway) Run(router *gin.Engine, routerPort string) error {
	err := http.ListenAndServe(routerPort, router)
	if err != nil {
		return err
	}
	return nil
}
