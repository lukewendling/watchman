package main

type events []event

// watchman event struct
type event struct {
	ID                string         `json:"id"`
	Name              string         `json:"name"`
	Keywords          keywordPairs   `json:"keywords"`
	Hashtags          []string       `json:"hashtags"`
	URLs              []string       `json:"urls"`
	ImageURLs         []string       `json:"image_urls"`
	Domains           []string       `json:"domains"`
	TopicMessageCount int            `json:"topic_message_count"`
	Locations         locations      `json:"location"`
	StartTimeMs       int64          `json:"start_time_ms"`
	EndTimeMs         int64          `json:"end_time_ms"`
	CampaignScores    campaignScores `json:"campaigns"`
}

// ex: [{'345678': 0.958139534883721}, {'123456': 0.958139534883721}]
type campaignScores []campaignScore
type campaignScore map[string]float64

// list of mixed-type pairs: [(crisis, 20.0)]
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
	words := make([]string, 0)
	for i, v := range kp {
		words[i] = v[0].(string)
	}
	return words
}

type location struct {
	GeoType string  `json:"geo_type" form:"geoType" validation:"required"`
	Coords  coords  `json:"coords"`
	Type    string  `json:"type"`
	Weight  float64 `json:"weight"`
	Label   string  `json:"label"`
}
type locations []location
type coords []map[string]float64

func (loc locations) Len() int {
	return len(loc)
}
func (loc locations) Swap(i, j int) {
	loc[i], loc[j] = loc[j], loc[i]
}
func (loc locations) Less(i, j int) bool {
	// desc by weight
	return loc[i].Weight > loc[j].Weight
}
