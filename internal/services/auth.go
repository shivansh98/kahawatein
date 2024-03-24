package services

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shivansh98/kahawatein/internal/adapter/cache"
	"github.com/shivansh98/kahawatein/internal/adapter/database"
	"github.com/shivansh98/kahawatein/internal/adapter/database/models"
	"github.com/shivansh98/kahawatein/internal/dto"
	. "github.com/shivansh98/kahawatein/utilities"
	"io/ioutil"
	"time"
)

func SignUp(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	req := dto.SignUpRequest{}
	err = json.Unmarshal(b, &req)
	if err != nil {
		c.AbortWithStatusJSON(404, `{"error":"Error unmarshaling request body"}`)
	}
	user := models.User{
		Username: req.Username,
		Password: req.Password,
		EmailID:  req.Email,
	}
	jwt, err := database.CreateUserProfile(c.Request.Context(), &user)
	if err != nil {
		Logger.Println("got an error in creating jwt token ", err)
		c.AbortWithStatusJSON(404, "failed to create user profile")
		return
	}
	c.SetCookie("token", jwt, int(time.Now().Add(5*time.Minute).Unix()), "", "localhost:8080", true, true)
	c.IndentedJSON(200, "success")
}

func SignIn(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	req := dto.SignUpRequest{}
	err = json.Unmarshal(b, &req)
	if err != nil {
		c.AbortWithStatusJSON(404, `{"error":"Error unmarshaling request body"}`)
		return
	}
	user := models.User{
		Username: req.Username,
		Password: req.Password,
		EmailID:  req.Email,
	}

	if !database.IsUserExists(c.Request.Context(), &user) {
		Logger.Println("user doesnt exists", err)
		c.AbortWithStatusJSON(404, "user doesnt exists")
		return
	}
	jwt, err := database.CreateJWT(&user)
	if err != nil {
		Logger.Println("failed to create jwt", err)
		c.AbortWithStatusJSON(404, "failed to create jwt")
		return
	}
	r := cache.GetRedisClient()
	_, err = r.Set(jwt, user.Username)
	if err != nil {
		Logger.Println("failed to set key to cache", err)
		c.AbortWithStatusJSON(404, "failed to set key to cache")
		return
	}
	c.SetCookie("token", jwt, int(time.Now().Add(10*time.Minute).Unix()), "", ":8080", true, true)
	c.IndentedJSON(200, "success")
}
