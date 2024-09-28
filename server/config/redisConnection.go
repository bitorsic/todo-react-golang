package config

import (
	"context"
	"fmt"
	"os"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis Connected")
}
