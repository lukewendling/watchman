package main

import (
	_ "fmt"
	"testing"
)

var sverige = location{
	Coords: coords{
		{
			"lat": 59.6749712,
			"lng": 14.5208584,
		},
	},
	GeoType: "point",
	Type:    "inferred point",
	Weight:  0.5,
	Label:   "Sverige",
}

var oslo = location{
	Coords: coords{
		{
			"lat": 59.911491,
			"lng": 10.757933,
		},
	},
	GeoType: "point",
	Type:    "inferred point",
	Weight:  0.8,
	Label:   "Oslo",
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
	CampaignScores: []campaignScore{
		{"000": 0.95},
		{"111": 0.5},
		{"222": 0.7},
	},
}

func Test_ToQCR_happy_path(t *testing.T) {

	qcrEvents := ToQCR(testEvent)

	if len(qcrEvents) != 2 {
		t.Error("should be 2 but was ", len(qcrEvents))
	}

	evt := qcrEvents[0]

	if evt["uid"] != "123" {
		t.Error("should be '123' but was ", evt["uid"])
	}

	kw := evt["keywords"].([]string)[0]
	if kw != "crisis" {
		t.Error("should be 'crisis' but was ", kw)
	}

	ht := evt["hashtags"].([]string)[0]
	if ht != "pakistan" {
		t.Error("should be 'pakistan' but was ", ht)
	}

	sd := evt["startDate"].(string)
	if sd != "2017-02-16 16:27:17 +0000 UTC" {
		t.Error("should not be ", sd)
	}

	cid := evt["campaignId"].(string)
	if cid != "000" {
		t.Error("should not be ", cid)
	}

	loc := evt["location"].(qcrLoc)
	co := loc["coordinates"].(qcrCoords)
	if co[0] != 10.757933 {
		t.Error("should not be ", co[0])
	}
}

func Test_ToQCR_empty_locations(t *testing.T) {

	testEvent.Locations = locations{}

	evt := ToQCR(testEvent)[0]

	loc := evt["location"].(qcrLoc)

	if len(loc) != 0 {
		t.Error("should have no location but was ", loc)
	}
}

func Test_ToQCR_nil_locations(t *testing.T) {

	testEvent.Locations = nil

	evt := ToQCR(testEvent)[0]

	loc := evt["location"].(qcrLoc)

	if len(loc) != 0 {
		t.Error("should have no location but was ", loc)
	}
}
