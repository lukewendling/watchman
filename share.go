package main

import (
	"log"
	"os"

	_ "fmt"

	"github.com/Shopify/sarama"
)

// ShareEvent is the main interface
func ShareEvent(evt *event) {
	kafkaURL := os.Getenv("KAFKA_URL")
	if kafkaURL == "" {
		kafkaURL = "172.17.0.1:9092"
	}

	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	if kafkaTopic == "" {
		kafkaTopic = "events"
	}

	ToQCR(*evt)

	producer, err := sarama.NewSyncProducer([]string{kafkaURL}, nil)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: kafkaTopic,
		Value: sarama.StringEncoder(evt.Name)}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}
