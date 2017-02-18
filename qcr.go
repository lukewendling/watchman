package main

import (
	"encoding/json"
	"fmt"
    "sort"
)

// Format is QCR structure
type Format struct {
    UID string `json:"id"`
    Label string 
    // StartDate string
    // EndDate string
    // Domains []string
    Hashtags []string `json:"hashtags"`
	Keywords [][]interface{} `json:"keywords"`
    // Urls []string
    // Photos []string
    // ImportanceScore int
    // TopicMessageCount int
    // CampaignID string
    // NewsEventsIDs []string
    // Location []interface{}
}

// ToQCR parses event
func ToQCR(data []byte) {
    // var parsed map[string]interface{}

    var formatted Format
    
    err := json.Unmarshal(data, &formatted)

    keywords := getKeywords(formatted.Keywords)

    var strSlice sort.StringSlice = keywords

    sort.Sort(sort.Reverse(strSlice)) 

    formatted.Label = Get(formatted.Hashtags, 0, "None")
    // formatted.Keywords = strSlice

    fmt.Println(err, formatted)
}

func getKeywords(vs [][]interface{}) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = v[0].(string)
	}
	return vsm
}

//keywords = map(iget(0), sorted(rec['keywords'], key=iget(1), reverse=True))
// l_rec.append({
//             'uid': rec['id'],
//             'label': rec['hashtags'][0] if len(rec['hashtags']) > 0 else 'None',
//             'startDate': datetime.fromtimestamp(rec['start_time_ms']/1000.0).isoformat(),
//             'endDate': datetime.fromtimestamp(rec['end_time_ms']/1000.0).isoformat(),
//             'domains': rec['domains'],
//             'hashtags': rec['hashtags'],
//             'keywords': keywords,
//             'urls': rec['urls'],
//             'photos': rec['image_urls'],
//             'importanceScore': camp[1],
//             'topicMessageCount':rec['topic_message_count'],
//             'campaignId': camp[0],
//             'newsEventIds': [],
//             'location': o_loc
//         })