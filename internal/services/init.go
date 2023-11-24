package services

import (
	"github.com/gin-gonic/gin"
	"github.com/shivansh98/kahawatein/internal/middlewear"
	. "github.com/shivansh98/kahawatein/utilities"
	"net/http"
)

func InitHTTPServer() *http.Server {
	Logger.Println("starting http server")

	r := gin.Default()
	// singup
	r.POST("/api/v1/signup", SignUp)
	// All auth based routes go here
	auth := r.Group("/api/v1/auth")
	auth.Use(middlewear.AuthMiddleWear)
	auth.GET("/feed", Feed)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}
	return server
}
