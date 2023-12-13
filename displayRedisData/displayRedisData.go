package displayRedisData

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/takehome/sendDataToKafka"
	"github.com/redis/go-redis/v9"
)

func DisplayRedisData(w http.ResponseWriter, r *http.Request) {
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Get an iterator to iterate over all keys using SCAN command
	iter := client.Scan(context.Background(), 0, "", 0).Iterator()

	// Map to store key-value pairs retrieved from Redis
	redisData := make(map[string]string)

	// Iterate through keys and fetch their values
	for iter.Next(context.Background()) {
		key := iter.Val()
		val, err := client.Get(context.Background(), key).Result()
		if err != nil {
			fmt.Printf("Error getting value for key '%s': %s\n", key, err)
			continue
		}
		redisData[key] = val
	}

	if err := iter.Err(); err != nil {
		fmt.Println("Error during iteration:", err)
	}

	// Convert Redis data to JSON
	jsonData, err := json.Marshal(redisData)
	if err != nil {
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
		return
	}

	// Set response headers and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

	go sendDataToKafka.SendDataToKafka(jsonData)

}
