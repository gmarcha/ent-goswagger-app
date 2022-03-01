package redis

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func Init() *redis.Client {

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
		DB:       0,
	})
	log.Println("Redis connected")

	_, err := rdb.Ping(context.Background()).Result()

	if err != nil {
		log.Fatalf("Failed to ping redis")
	}

	return rdb
}
