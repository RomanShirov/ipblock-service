package backend

import (
	"github.com/go-redis/redis/v8"
	"time"
)

func InitRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:         "localhost:6379",
		DialTimeout:  1 * time.Second,
		ReadTimeout:  50 * time.Millisecond,
		WriteTimeout: 300 * time.Millisecond,
		PoolSize:     1500,
		PoolTimeout:  300 * time.Millisecond,
	}
}
