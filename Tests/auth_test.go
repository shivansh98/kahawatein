package tests

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shivansh98/kahawatein/auth"
)

func AuthTest(t *testing.T) {
	router := gin.Default()
	router.GET("/sigin", auth.CheckAuth)
	router.Run("localhost:8080")

}
