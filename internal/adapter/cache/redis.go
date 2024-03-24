package cache

import (
	"time"

	"github.com/go-redis/redis/v8"
)

const DefaultTTL = 600

type Redis struct {
	Client redis.UniversalClient
}

var RedisClient *Redis

func InitRedisClient(opt *redis.Options) *Redis {
	if RedisClient != nil {
		return RedisClient
	}
	cl := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:     []string{opt.Addr},
		Username:  opt.Username,
		Password:  opt.Password,
		DB:        opt.DB,
		PoolSize:  opt.PoolSize,
		TLSConfig: opt.TLSConfig,
	})
	RedisClient = &Redis{
		Client: cl,
	}
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
