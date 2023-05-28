package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/shivansh98/kahawatein/internal/middlewear"
	. "github.com/shivansh98/kahawatein/utilities"
	"net/http"
)

func InitHTTPServer(ctx context.Context) {
	defer func() {
		ctx.Done()
		Logger.Println("cancel func called")
	}()
	Logger.Println("starting http server")

	r := gin.Default()
	// singup
	r.POST("/api/v1/signup", SignUp)
	// All auth based routes go here
	auth := r.Group("/api/v1/auth")
	auth.Use(middlewear.AuthMiddleWear)
	auth.GET("/feed", Feed)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		Logger.Println("error occured ", err)
	}
}
