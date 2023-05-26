package services

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shivansh98/kahawatein/adapter/database"
	"github.com/shivansh98/kahawatein/adapter/database/models"
	"github.com/shivansh98/kahawatein/dto"
)

func InitHTTPServer() {
	r := gin.Default()
	r.POST("/api/v1/signup", SignUp)

}

func SignUp(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	req := dto.SignUpRequest{}
	err = json.Unmarshal(b, &req)
	if err != nil {
		c.AbortWithStatusJSON(404, "Error unmarshaling request body")
	}
	user := models.User{
		Username: req.Username,
		Password: req.Password,
		EmailID:  req.Email,
	}
	jwt, err := database.CreateUserProfile(c.Request.Context(), &user)
	if err != nil {
		c.AbortWithStatusJSON(404, "failed to create user profile")
		return
	}
	c.SetCookie("token", jwt, int(time.Now().Add(5*time.Minute).Unix()), "", "localhost:8080", true, true)
	c.IndentedJSON(200, "success")
}
