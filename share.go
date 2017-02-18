package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type events []event

type event struct {
	ID   string `json:"id"`
	Name string `json:"name"`

    Hashtags []string `json:"hashtags"`
	Keywords [][]interface{} `json:"keywords"`
}

// ShareEvent is the main interface
func ShareEvent(evt *event) {
	kafkaURL := os.Getenv("KAFKA_URL")
	if kafkaURL == "" {
		kafkaURL = "172.17.0.1:9092"
	}

	data, err := json.Marshal(evt)

	ToQCR(data)

	producer, err := sarama.NewSyncProducer([]string{kafkaURL}, nil)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: "events",
		Value: sarama.StringEncoder(evt.Name)}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}
