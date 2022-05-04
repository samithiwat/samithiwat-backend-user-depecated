package database

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/samithiwat/samithiwat-backend-user/src/config"
)

func InitRedisConnect(conf *config.Redis) (cache *redis.Client, err error) {
	cache = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		DB:   0,
	})

	return
}
