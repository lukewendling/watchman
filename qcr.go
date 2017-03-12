package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/Shopify/sarama"
)

// QCR type
type QCR struct{}

// Receive interface implementation
func (q *QCR) Receive(evts events) {
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
		qcrEvents := convert(evt)

		for _, qcrEvt := range qcrEvents {
			msg := &sarama.ProducerMessage{
				Topic: kafkaTopic,
				Value: sarama.StringEncoder(stringifyGenericMap(qcrEvt)),
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

// convert event into multiple QCR events, 1 per campaign
func convert(evt event) []map[string]interface{} {
	fmt.Println("converting event ", evt.ID)

	qcrEvent := map[string]interface{}{}

	sort.Sort(evt.Keywords)
	sort.Sort(evt.Locations)

	qcrEvent["keywords"] = evt.Keywords.keywords()
	qcrEvent["uid"] = evt.ID
	qcrEvent["label"] = Get(evt.Hashtags, 0, "None")
	qcrEvent["hashtags"] = evt.Hashtags
	qcrEvent["photos"] = evt.ImageURLs
	qcrEvent["urls"] = evt.URLs
	qcrEvent["domains"] = evt.Domains
	qcrEvent["topicMessageCount"] = evt.TopicMessageCount
	qcrEvent["newsEventIds"] = []string{}
	qcrEvent["location"] = toLocation(evt.Locations)
	qcrEvent["startDate"] = millisToTime(evt.StartTimeMs)
	qcrEvent["endDate"] = millisToTime(evt.EndTimeMs)

	qcrEvents := []map[string]interface{}{}

	if len(evt.CampaignScores) == 0 {
		fmt.Println("no campaigns found. return empty set.")
		return qcrEvents
	}

	for _, c := range evt.CampaignScores {
		copyEvent := map[string]interface{}{}
		// make a copy
		for k, v := range qcrEvent {
			copyEvent[k] = v
		}
		for k, v := range c {
			copyEvent["campaignId"] = k
			copyEvent["importanceScore"] = v
		}
		if copyEvent["importanceScore"].(float64) >= 0.7 {
			qcrEvents = append(qcrEvents, copyEvent)
		}
	}

	return qcrEvents
}

func millisToTime(millis int64) string {
	return time.Unix(millis/1000, 0).UTC().String()
}

func toLocation(locs locations) qcrLoc {
	if len(locs) == 0 {
		return qcrLoc{}
	}
	loc := locs[0]
	c := loc.Coords
	latLng := qcrCoords{c[0]["lng"], c[0]["lat"]}
	return qcrLoc{
		"type":        "Point",
		"coordinates": latLng,
	}
}

type qcrLoc map[string]interface{}
type qcrCoords []float64
