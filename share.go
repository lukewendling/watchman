package main

import (
	"encoding/json"
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

	qcrEvt := ToQCR(*evt)

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	brokers := []string{kafkaURL}
	producer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: kafkaTopic,
		Value: sarama.StringEncoder(stringifyEvent(qcrEvt)),
	}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}

func stringifyEvent(evt map[string]interface{}) string {
	b, err := json.Marshal(evt)
	if err != nil {
		log.Fatalln(err)
	}
	return string(b)
}
