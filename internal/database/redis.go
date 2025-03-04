package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	// Retrieve Redis env
	containerName := os.Getenv("REDIS_CONTAINER_NAME")
	port := os.Getenv("REDIS_DOCKER_PORT")

	// Connect
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", containerName, port),
		Password: "",
		DB:       0,
	})

	// Test connection
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect Redis: %v", err)
	}

	return rdb
}
