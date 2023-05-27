package services

import (
	"encoding/json"
	. "github.com/shivansh98/kahawatein/utilities"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shivansh98/kahawatein/adapter/database"
	"github.com/shivansh98/kahawatein/adapter/database/models"
	"github.com/shivansh98/kahawatein/dto"
)

func InitHTTPServer() {
	defer func() {
		if r := recover(); r != nil {
			Logger.Println("We got a panic in InitHTTPServer")
		}
	}()
	Logger.Println("starting http server")
	r := gin.Default()
	r.POST("/api/v1/signup", SignUp)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		Logger.Println("error occured ", err)
	}
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
