package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	startTime := flag.String("start-time-ms", "", "start time in millis")
	endTime := flag.String("end-time-ms", "", "end time in millis")

	flag.Parse()

	apiRoot := os.Getenv("API_ROOT")
	if apiRoot == "" {
		apiRoot = "http://172.17.0.1:3000/api"
	}

	qs := Between("created", *startTime, *endTime)

	res, err := http.Get(apiRoot + "/events" + qs)

	if err != nil {
		log.Fatal(fmt.Println(err))
	}

	watchmanEvents := make(events, 0)

	json.NewDecoder(res.Body).Decode(&watchmanEvents)
	fmt.Println("events:", len(watchmanEvents))

	ShareEvents(watchmanEvents)
}
