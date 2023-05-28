package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Feed(c *gin.Context) {
	var user any
	var exists bool
	if user, exists = c.Get("username"); !exists {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusOK, map[string]string{
		"status": "ok",
		"feed":   "here will be feed .definitely",
		"user":   "thanks for logging in user" + user.(string),
	})
}
