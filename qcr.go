package main

import (
	_ "fmt"
	"sort"
	"time"
)

// ToQCR converts event into QCR format
func ToQCR(evt event) map[string]interface{} {
	qcrEvent := make(map[string]interface{})

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

	// fmt.Println(qcrEvent)
	return qcrEvent
}

func millisToTime(millis int64) string {
	return time.Unix(millis/1000, 0).UTC().String()
}

func toLocation(locs locations) qcrLoc {
	if len(locs) == 0 {
		return qcrLoc{}
	}
	loc := locs[0]
	c := loc["coords"].(coords)[0]
	latLng := qcrCoords{c["lng"], c["lat"]}
	return qcrLoc{
		"type":        "Point",
		"coordinates": latLng,
	}
}

type qcrLoc map[string]interface{}
type qcrCoords []float64
