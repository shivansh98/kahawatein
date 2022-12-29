package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shivansh98/kahawatein/auth"
)

func main() {
	router := gin.Default()
	router.GET("/hellow", welcome)

	router.GET("/sigin", auth.CheckAuth)
	router.Run("localhost:8080")
}

type User struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"pass,omitempty"`
}

var a = User{
	Name:     "Shivansh",
	Password: "shivalfaz",
}

func welcome(c *gin.Context) {

	c.IndentedJSON(200, a)
}
