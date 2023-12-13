package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Flush all keys and their content in Redis
	ctx := context.Background()
	statusCmd := rdb.FlushAll(ctx)
	if err := statusCmd.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted all keys and their content in Redis.")
}
