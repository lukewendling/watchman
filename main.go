package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Sotera/go_watchman/loogo"
)

func main() {
	startTime := flag.String("start-time-ms", "", "start time in millis")
	endTime := flag.String("end-time-ms", "", "end time in millis")

	flag.Parse()

	apiRoot := os.Getenv("API_ROOT")
	if apiRoot == "" {
		apiRoot = "http://172.17.0.1:3000/api"
	}
	apiRoot = strings.TrimRight(apiRoot, "/")

	parser := &loogo.HTTPRequestParser{Client: &loogo.HTTPClient{}}
	p1 := loogo.QueryParam{
		QueryType: "between",
		Field:     "created",
		Values:    []string{*startTime, *endTime},
	}
	params := loogo.QueryParams{
		p1,
	}

	watchmanEvents := make(events, 0)

	err := parser.NewRequest(
		loogo.NewRequestParams{URL: apiRoot + "/events", Params: params},
		watchmanEvents)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("events:", len(watchmanEvents))

	ShareEvents(watchmanEvents)
}
