package utils

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func RedisConnect() (rdb *redis.Client) {

	var (
		redisServer string
		port        string
		password    string
	)

	env := os.Getenv("GIN_MODE")

	if env == "release" {
		redisServer = os.Getenv("RedisUrl")
		port = os.Getenv("RedisPort")
		password = os.Getenv("RedisPass")
	} else {
		redisServer = "localhost"
		port = "6379"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisServer + ":" + port,
		Password: password,
		DB:       0, // use default DB
	})

	return
}
