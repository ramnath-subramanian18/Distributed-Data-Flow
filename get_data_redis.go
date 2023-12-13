package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/takehome/sendDataToKafka"
	"github.com/redis/go-redis/v9"
)

// func send_data_kafka(msg []byte) {
// 	config := sarama.NewConfig()
// 	config.Producer.RequiredAcks = sarama.WaitForAll
// 	config.Producer.Retry.Max = 5
// 	config.Producer.Return.Successes = true

// 	// Define Kafka broker addresses
// 	brokers := []string{"localhost:9092"} // Replace with your Kafka broker addresses

// 	// Create a producer
// 	producer, err := sarama.NewSyncProducer(brokers, config)
// 	if err != nil {
// 		log.Fatalf("Failed to create producer: %s", err)
// 	}
// 	defer func() {
// 		if err := producer.Close(); err != nil {
// 			log.Fatalf("Failed to close producer: %s", err)
// 		}
// 	}()

// 	// Produce a message to a topic
// 	topic := "final_topic" // Replace with your Kafka topic
// 	message := &sarama.ProducerMessage{
// 		Topic: topic,
// 		Value: sarama.ByteEncoder(msg),
// 	}

// 	// Send message to Kafka
// 	partition, offset, err := producer.SendMessage(message)
// 	if err != nil {
// 		log.Fatalf("Failed to send message: %s", err)
// 	}

// 	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)

// }

func displayRedisData(w http.ResponseWriter, r *http.Request) {
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

	go sendDataToKafka.SendDataToKafka(jsonData, "final")

	// ctx := context.Background()
	// statusCmd := client.FlushAll(ctx)
	// if err := statusCmd.Err(); err != nil {
	// 	fmt.Printf("Error deleting the data")
	// }

	// fmt.Println("Deleted all keys and their content in Redis.")

}

func main() {
	// Set up an endpoint to display Redis data
	http.HandleFunc("/displayRedisData", displayRedisData)

	// Start HTTP server
	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
