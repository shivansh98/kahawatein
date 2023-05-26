package database

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/shivansh98/kahawatein/adapter/cache"
	"github.com/shivansh98/kahawatein/adapter/database/models"
	"github.com/shivansh98/kahawatein/constant"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateUserProfile(ctx context.Context, u *models.User) (string, error) {
	if isUserExists(ctx, u) {
		return "", fmt.Errorf("user already exists")
	}
	var err error
	u.JWT, err = createJWT(u)
	if err != nil {
		return "", err
	}
	client := GetConnection(ctx)
	resp, err := client.Database(viper.GetString("MONGO_DATABASE")).Collection(string(constant.COLLECTION_USER)).InsertOne(ctx, u)
	if err != nil {
		return "", err
	}
	if resp.InsertedID == nil {
		return "", fmt.Errorf("failed to insert document in DB")
	}

	r := cache.GetRedisClient()
	if _, err = r.Set(u.Username, u.JWT); err != nil {
		log.Default().Println("error in inserting the key in redis")
	}

	return u.JWT, nil
}

func isUserExists(ctx context.Context, u *models.User) bool {
	res := GetConnection(ctx).Database(viper.GetString("MONGO_DATABASE")).Collection(string(constant.COLLECTION_USER)).FindOne(ctx, u)
	if res.Err() != nil {
		return true //taking it as safety to not make another entry
	}
	var resp models.User
	err := res.Decode(&resp)

	if err != nil && resp.EmailID == u.EmailID {
		return true
	}
	return false
}

func createJWT(u *models.User) (string, error) {
	claims := Claims{
		Username: u.Username,
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(10 * time.Minute))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	sign, err := token.SignedString(viper.GetString("JWT_KEY"))
	if err != nil {
		return "", err
	}
	return sign, nil
}
