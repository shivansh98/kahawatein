package middlewear

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shivansh98/kahawatein/internal/adapter/cache"
	. "github.com/shivansh98/kahawatein/utilities"
	"net/http"
)

func AuthMiddleWear(c *gin.Context) {
	req := c.Request
	ck, err := req.Cookie("token")
	if err != nil {
		Logger.Println("auth middlewear got an error ", "error in fetching cookie token error:", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	val := cache.GetRedisClient().Get(ck.Value)
	if val == "" {
		Logger.Println("auth middlewear got an error ", "user not authorized:")
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("user not authorized"))
		return
	}

	c.Set("username", val) // setting username in request context
	c.Next()
}
