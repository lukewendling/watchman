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
	StartTimeMs:       1487262437000,
	EndTimeMs:         1487262736999,
}

func Test_ToQCR_happy_path(t *testing.T) {

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

	sd := got["startDate"].(string)
	if sd != "2017-02-16 16:27:17 +0000 UTC" {
		t.Error("should not be ", sd)
	}

	loc := got["location"].(qcrLoc)
	c := loc["coordinates"].(qcrCoords)
	if c[0] != 10.757933 {
		t.Error("should not be ", c[0])
	}
}

func Test_ToQCR_empty_locations(t *testing.T) {

	testEvent.Locations = locations{}

	got := ToQCR(testEvent)

	loc := got["location"].(qcrLoc)

	if len(loc) != 0 {
		t.Error("should have no location but was ", loc)
	}
}

func Test_ToQCR_nil_locations(t *testing.T) {

	testEvent.Locations = nil

	got := ToQCR(testEvent)

	loc := got["location"].(qcrLoc)

	if len(loc) != 0 {
		t.Error("should have no location but was ", loc)
	}
}
