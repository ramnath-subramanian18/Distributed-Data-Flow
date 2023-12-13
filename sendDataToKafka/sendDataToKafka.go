package sendDataToKafka

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func SendDataToKafka(msg []byte, topic string) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// config := sarama.NewConfig()
	// producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	brokers := []string{"localhost:9092"} // Replace with your Kafka broker addresses

	// Create a producer
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Failed to close Kafka producer: %s", err)
		}
	}()

	// topic := topic // Replace with your Kafka topic
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	}

	// Send message to Kafka
	_, _, err = producer.SendMessage(message)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %s", err)
	}
	fmt.Printf("data sent into kafka")
}
