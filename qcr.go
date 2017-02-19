package main

import (
	"fmt"
	"sort"
)

// ToQCR parses event to QCR format
func ToQCR(evt event) map[string]interface{} {
	qcrEvent := make(map[string]interface{})

	keywords := GetKeywords(evt.Keywords)

	var strSlice sort.StringSlice = keywords

	qcrEvent["keywords"] = strSlice

	sort.Sort(sort.Reverse(strSlice))

	qcrEvent["uid"] = evt.ID
	qcrEvent["label"] = Get(evt.Hashtags, 0, "None")
	qcrEvent["hashtags"] = evt.Hashtags
	qcrEvent["photos"] = evt.ImageURLs
	qcrEvent["urls"] = evt.URLs
	qcrEvent["domains"] = evt.Domains
	qcrEvent["topicMessageCount"] = evt.TopicMessageCount

	fmt.Println(qcrEvent)
	return qcrEvent
}

// GetKeywords takes keywords from keyword->count pairs
func GetKeywords(vs [][]interface{}) []string {
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
