package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/n7down/microservices/internal/messages"
	messagesPB "github.com/n7down/microservices/internal/messages/pb"

	"github.com/n7down/microservices/internal/users"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var (
	identityKey  = "id"
	gatewayPort  = "8080"
	messagesPort = "8081"
)

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	v1 := router.Group("/api/v1")

	auth := v1.Group("/auth")
	auth.GET("/refresh", authMiddleware.RefreshHandler)
	auth.POST("/login", authMiddleware.LoginHandler)

	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	helloClient := messagesPB.NewHelloServiceClient(conn)

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
