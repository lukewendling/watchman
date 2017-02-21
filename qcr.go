package main

import (
	"fmt"
	"sort"
	"time"
)

// ToQCR converts event into multiple QCR events, 1 per campaign
func ToQCR(evt event) []map[string]interface{} {
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

	campaignEvents := []map[string]interface{}{}

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
			campaignEvents = append(campaignEvents, copyEvent)
		}
	}

	return campaignEvents
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
