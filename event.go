package main


type events []event

type event struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Keywords keywordPairs `json:"keywords"`
    Hashtags []string `json:"hashtags"`
	URLs []string `json:"urls"`
	ImageURLs []string `json:"image_urls"`
	Domains []string `json:"domains"`
	TopicMessageCount int `json:"topic_message_count"`
}

// sortable pairs: [word string, count int]
type keywordPairs [][]interface{}

func (kp keywordPairs) Len() int {
    return len(kp)
}
func (kp keywordPairs) Swap(i, j int) {
    kp[i], kp[j] = kp[j], kp[i]
}
func (kp keywordPairs) Less(i, j int) bool {
	// desc by count: element index 1
    return kp[i][1].(int) > kp[j][1].(int)
}

// extract keywords from pairs
func (kp keywordPairs) keywords() []string {
	words := make([]string, len(kp))
	for i, v := range kp {
		words[i] = v[0].(string)
	}
	return words
}