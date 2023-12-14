package sendDataToKafka

//data sent into kafka
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

	brokers := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Failed to close Kafka producer: %s", err)
		}
	}()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	}

	_, _, err = producer.SendMessage(message)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %s", err)
	}
	fmt.Printf("data sent into kafka")
}
