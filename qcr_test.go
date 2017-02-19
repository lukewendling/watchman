package main

import (
	_ "fmt"
	"testing"
)

var sverige = location{
	"coords": coords{
		{
			"lat": 59.6749712,
			"lng": 14.5208584,
		},
	},
	"geo_type": "point",
	"type":     "inferred point",
	"weight":   0.5,
	"label":    "Sverige",
}

var oslo = location{
	"coords": coords{
		{
			"lat": 59.911491,
			"lng": 10.757933,
		},
	},
	"geo_type": "point",
	"type":     "inferred point",
	"weight":   0.8,
	"label":    "Oslo",
}

var testEvent = event{
	ID:                "123",
	Name:              "name",
	Keywords:          keywordPairs{[]interface{}{"attack", 10.0}, []interface{}{"crisis", 20.0}},
	Hashtags:          []string{"pakistan", "daesh", "isis"},
	URLs:              []string{"https://t.co/wi29dkomyd", "https://t.co/njm1iclohh"},
	ImageURLs:         []string{"https://pbs.twimg.com/media/C4xo5cEWEAAksFH.jpg:large", "https://pbs.twimg.com/media/C4zMvyLUYAAfeZT.jpg:large"},
	TopicMessageCount: 133,
	Locations:         locations{sverige, oslo},
}

func TestToQCR(t *testing.T) {

	got := ToQCR(testEvent)

	if got["uid"] != "123" {
		t.Error("should be '123' but was ", got["uid"])
	}

	k := got["keywords"].([]string)[0]
	if k != "crisis" {
		t.Error("should be 'crisis' but was ", k)
	}

	h := got["hashtags"].([]string)[0]
	if h != "pakistan" {
		t.Error("should be 'pakistan' but was ", h)
	}

	loc := got["location"].(location)
	if loc["label"] != "Oslo" {
		t.Error("should be 'Oslo' but was ", loc["label"])
	}
}
