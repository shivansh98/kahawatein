package bootstrap

import (
	"github.com/go-redis/redis/v8"
	"github.com/gookit/ini/v2/dotenv"
	"github.com/shivansh98/kahawatein/adapter/cache"
	"github.com/shivansh98/kahawatein/services"
	"github.com/shivansh98/kahawatein/utilities"
	"github.com/spf13/viper"
)

func InitServices() {
	err := dotenv.Load("./", ".env")
	if err != nil {
		utilities.CallPanic(err)
	}
	viper.AutomaticEnv()
	rp := redis.Options{
		Addr: viper.GetString("REDIS_ADDRESS"),
		DB:   viper.GetInt("REDIS_DATABASE"),
	}
	cache.InitRedisClient(&rp)
	utilities.InitLogger()
	services.InitHTTPServer()

}
