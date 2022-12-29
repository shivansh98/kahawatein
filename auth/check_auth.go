package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAuth(c *gin.Context) {
	fmt.Print("checked in here")
	c.IndentedJSON(http.StatusOK, "|Hello")

}
