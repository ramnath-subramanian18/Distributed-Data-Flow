package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/takehome/sendDataToKafka"
	"github.com/redis/go-redis/v9"
)

func displayRedisData(w http.ResponseWriter, r *http.Request) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	iter := client.Scan(context.Background(), 0, "", 0).Iterator()

	redisData := make(map[string]string)

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

	jsonData, err := json.Marshal(redisData)
	if err != nil {
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//display the data using GET request
	w.Write(jsonData)
	//send the data into new kafka topic
	//goroutine-3,
	go sendDataToKafka.SendDataToKafka(jsonData, "final")
}

func main() {
	//GET request displayRedisData
	http.HandleFunc("/displayRedisData", displayRedisData)

	fmt.Println("Server listening on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
