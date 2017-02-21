package main

import (
	"encoding/json"
	"log"
	"os"

	_ "fmt"

	"github.com/Shopify/sarama"
)

// ShareEvents publishes events to QCR msg queue
func ShareEvents(evts events) {
	kafkaURL := os.Getenv("KAFKA_URL")
	if kafkaURL == "" {
		kafkaURL = "172.17.0.1:9092"
	}

	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	if kafkaTopic == "" {
		kafkaTopic = "events"
	}

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

	for _, evt := range evts {
		qcrEvents := ToQCR(evt)

		for _, qcrEvt := range qcrEvents {
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
	}
}

func stringifyEvent(evt map[string]interface{}) string {
	b, err := json.Marshal(evt)
	if err != nil {
		log.Fatalln(err)
	}
	return string(b)
}
