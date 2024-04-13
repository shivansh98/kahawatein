package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/shivansh98/kahawatein/internal/middlewear"
	"github.com/shivansh98/kahawatein/internal/services"
	. "github.com/shivansh98/kahawatein/utilities"
	"net/http"
)

func InitHTTPServer() *http.Server {
	Logger.Println("starting http server")

	r := gin.Default()
	r.Use(gin.CustomRecovery(PanicHandler))
	// singup
	r.POST("/api/v1/signup", services.SignUp)
	r.POST("/api/v1/signin", services.SignIn)
	r.GET("/", services.Home)
	// All auth based routes go here
	auth := r.Group("/api/v1/auth")
	auth.Use(middlewear.AuthMiddleWear)
	auth.GET("/feed", services.Feed)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}
	return server
}

func PanicHandler(c *gin.Context, err any) {
	Logger.Println(err)
}
