package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	expireTime  = time.Hour * 6
	rdToken     *redis.Client
	memberToken *redis.Client
	ctx         = context.Background()
)

func InitRedis() {
	rdToken = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
	memberToken = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       1,
	})
}
