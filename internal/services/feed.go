package services

import (
	"github.com/gin-gonic/gin"
	external_service "github.com/shivansh98/kahawatein/internal/services/external-service"
	"net/http"
)

type FeedReq struct {
	Query string `json:"query" form:"query"`
}

func Feed(c *gin.Context) {
	fdRq := FeedReq{}
	err := c.BindQuery(&fdRq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	resp := external_service.SearchUnsplash(c.Request.Context(), fdRq.Query)
	c.IndentedJSON(http.StatusOK, map[string]interface{}{
		"response": resp,
	})
}
