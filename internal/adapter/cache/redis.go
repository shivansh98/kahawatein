package cache

import (
	"time"

	"github.com/go-redis/redis/v8"
)

const DefaultTTL = 600

type Redis struct {
	Client *redis.Client
}

var RedisClient *Redis

func InitRedisClient(rp *redis.Options) *Redis {
	if RedisClient != nil {
		return RedisClient
	}
	RedisClient = &Redis{Client: redis.NewClient(rp)}
	return RedisClient
}

func GetRedisClient() *Redis {
	return RedisClient
}

func (r *Redis) Set(key, value string) (string, error) {
	return r.Client.Set(r.Client.Context(), key, interface{}(value), DefaultTTL*time.Second).Result()
}

func (r *Redis) Get(key string) (value string) {
	val := r.Client.Get(r.Client.Context(), key)
	return val.Val()
}
