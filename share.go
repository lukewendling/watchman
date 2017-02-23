package main

import (
	"encoding/json"
	"log"

	_ "fmt"
)

// EventReceiver is interface for all receivers
type EventReceiver interface {
	Receive(events)
}

// ShareEvents sends events to specified receivers
func ShareEvents(evts events) {
	receivers := []EventReceiver{
		&QCR{},
	}
	for _, rec := range receivers {
		rec.Receive(evts)
	}
}

// marshal to json and stringify it, prolly for transmission
func stringifyGenericMap(evt map[string]interface{}) string {
	b, err := json.Marshal(evt)
	if err != nil {
		log.Fatalln(err)
	}
	return string(b)
}
