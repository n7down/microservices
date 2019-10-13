package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/n7down/microservices/internal/greeter"
	"github.com/n7down/microservices/internal/greeter/pb"

	"github.com/n7down/microservices/internal/users"
)

var (
	gatewayPort = "8080"
)

// LoginRequest represents a login json request.
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse defines a login response
type LoginResponse struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Token     string `json:"token"`
	Expires   string `json:"expires"`
}

func AuthLogin(in LoginRequest) (*LoginResponse, error) {
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

	if in.Username == "admin" && in.Password == "admin" {
		return &LoginResponse{
			FirstName: "Leslie",
			LastName:  "Knope",
			Username:  "leslie2020",
		}, nil
	}

	return &LoginResponse{}, jwt.ErrFailedAuthentication
}

var loginResponse *LoginResponse

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "production",
		Key:              []byte(os.Getenv("AUTH_KEY")),
		Timeout:          time.Duration(24) * time.Hour,
		MaxRefresh:       time.Duration(24) * time.Hour,
		IdentityKey:      "id",
		SigningAlgorithm: "HS256",
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var in LoginRequest
			if err := c.ShouldBind(&in); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			authLoginResponse := &LoginResponse{}
			authLoginResponse, err := AuthLogin(in)
			loginResponse = authLoginResponse

			if err != nil {
				return "", err
			}

			return authLoginResponse, nil
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

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	v1 := router.Group("/api/v1")

	auth := v1.Group("/auth")
	auth.GET("/refresh", authMiddleware.RefreshHandler)
	auth.POST("/login", authMiddleware.LoginHandler)

	// FIXME: move to somewhere else
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	helloClient := greeter_pb.NewHelloServiceClient(conn)

	messagesGroup := v1.Group("/messages")
	//internal.Use(authMiddleware.MiddlewareFunc())
	{
		m := messages.NewMessages(helloClient)
		messagesGroup.GET("/hello", m.HelloHandler)
	}

	usersGroup := v1.Group("/users")
	usersGroup.Use(authMiddleware.MiddlewareFunc())
	{
		u := users.NewUsers()
		usersGroup.POST("/create", u.CreateHandler)
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

	routerPort := fmt.Sprintf(":%s", gatewayPort)
	fmt.Printf("Listening on port: %s\n", gatewayPort)
	if err := http.ListenAndServe(routerPort, router); err != nil {
		log.Fatal(err)
	}
}
