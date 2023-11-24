package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/shivansh98/kahawatein/internal/services"
	. "github.com/shivansh98/kahawatein/utilities"
	"net/http"
)

func InitHTTPServer() *http.Server {
	Logger.Println("starting http server")

	r := gin.Default()
	// singup
	r.POST("/api/v1/signup", services.SignUp)
	// All auth based routes go here
	auth := r.Group("/api/v1/auth")
	//auth.Use(middlewear.AuthMiddleWear) TODO change this
	auth.POST("/feed", services.Feed)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}
	return server
}
