package redis

import (
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	expireTime = time.Hour * 1
	rdb        *redis.Client
)

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
}
