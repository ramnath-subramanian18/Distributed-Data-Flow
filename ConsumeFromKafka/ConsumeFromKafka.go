package ConsumeFromKafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/redis/go-redis/v9"
)

func ConsumeFromKafka() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalf("Error closing Kafka consumer: %v", err)
		}
	}()

	topic := "your_topic" // Replace with your Kafka topic
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating Kafka partition consumer: %v", err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalf("Error closing Kafka partition consumer: %v", err)
		}
	}()

	fmt.Println("Kafka consumer started. Waiting for messages...")

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,
	})

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Received message from Kafka: Value: %s, Offset: %d\n", string(msg.Value), msg.Offset)

			// Store message in Redis
			timestamp := time.Now().UnixNano()
			err := rdb.Set(ctx, fmt.Sprintf("%d", timestamp), string(msg.Value), 0).Err()
			// err := rdb.Set(ctx, timestamp, string(msg.Value), 0).Err()
			if err != nil {
				fmt.Println("Error setting key in Redis:", err)
			} else {
				fmt.Println("Key 'myKey' set successfully in Redis.")
			}
		case err := <-partitionConsumer.Errors():
			log.Printf("Error consuming from Kafka: %v", err)
		}
	}
}
