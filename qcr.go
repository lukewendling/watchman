package main

import (
	_ "fmt"
	"sort"
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

	// fmt.Println(qcrEvent)
	return qcrEvent
}

func toLocation(locs locations) qcrLoc {
    if len(locs) == 0 {
        return qcrLoc{}
    }
    loc := locs[0]
    c := loc["coords"].(coords)[0]
    latLng := qcrCoords{c["lng"], c["lat"]}
    return qcrLoc{
        "type": "Point",
        "coordinates": latLng,
    }
}

type qcrLoc map[string]interface{} 
type qcrCoords []float64

