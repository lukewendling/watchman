package main

import "testing"
import "sort"

func TestToQcr(t *testing.T) {
	testEvent := event{
		ID:                "123",
		Name:              "name",
		Keywords:          [][]interface{}{[]interface{}{"attack", 10}, []interface{}{"crisis", 20}},
		Hashtags:          []string{"pakistan", "daesh", "isis"},
		URLs:              []string{"https://t.co/wi29dkomyd", "https://t.co/njm1iclohh"},
		ImageURLs:         []string{"https://pbs.twimg.com/media/C4xo5cEWEAAksFH.jpg:large", "https://pbs.twimg.com/media/C4zMvyLUYAAfeZT.jpg:large"},
		TopicMessageCount: 133,
	}

	res := ToQCR(testEvent)

	if res["uid"] != "123" {
		t.Error("should be '123' but was ", res["uid"])
	}

	k := res["keywords"].(sort.StringSlice)[0]
	if k != "crisis" {
		t.Error("should be 'crisis' but was ", k)
	}
}

func TestGetKeywords(t *testing.T) {
	res := GetKeywords([][]interface{}{[]interface{}{"attack", 10}, []interface{}{"crisis", 20}})

	if res[0] != "attack" {
		t.Error("should be 'attack' but was ", res[0])
	}
}
