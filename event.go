package main

type events []event

// watchman event struct
type event struct {
	ID                string       `json:"id"`
	Name              string       `json:"name"`
	Keywords          keywordPairs `json:"keywords"`
	Hashtags          []string     `json:"hashtags"`
	URLs              []string     `json:"urls"`
	ImageURLs         []string     `json:"image_urls"`
	Domains           []string     `json:"domains"`
	TopicMessageCount int          `json:"topic_message_count"`
	Locations         locations    `json:"location"`
	StartTimeMs       int64        `json:"start_time_ms"`
	EndTimeMs         int64        `json:"end_time_ms"`
}

// pairs: [word string, count int]
type keywordPairs [][]interface{}

func (kp keywordPairs) Len() int {
	return len(kp)
}
func (kp keywordPairs) Swap(i, j int) {
	kp[i], kp[j] = kp[j], kp[i]
}
func (kp keywordPairs) Less(i, j int) bool {
	// desc by count: element index 1
	return kp[i][1].(float64) > kp[j][1].(float64)
}

// extract keywords from pairs
func (kp keywordPairs) keywords() []string {
	words := make([]string, len(kp))
	for i, v := range kp {
		words[i] = v[0].(string)
	}
	return words
}

type locations []location
type location map[string]interface{}
type coords []map[string]float64

func (loc locations) Len() int {
	return len(loc)
}
func (loc locations) Swap(i, j int) {
	loc[i], loc[j] = loc[j], loc[i]
}
func (loc locations) Less(i, j int) bool {
	// desc by weight
	return loc[i]["weight"].(float64) > loc[j]["weight"].(float64)
}
