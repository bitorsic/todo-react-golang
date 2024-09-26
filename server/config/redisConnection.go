package config

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func RedisConnect() {
	uri := os.Getenv("REDIS_URI")
	opt, err := redis.ParseURL(uri)
	if err != nil {
		panic(err)
	}

	RedisClient = redis.NewClient(opt)

	_, err = RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis Connected")
}
