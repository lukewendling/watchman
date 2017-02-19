package main

import (
	_ "fmt"
	"testing"
)

var testEvent = event{
	ID:                "123",
	Name:              "name",
	Keywords:          [][]interface{}{[]interface{}{"attack", 10}, []interface{}{"crisis", 20}},
	Hashtags:          []string{"pakistan", "daesh", "isis"},
	URLs:              []string{"https://t.co/wi29dkomyd", "https://t.co/njm1iclohh"},
	ImageURLs:         []string{"https://pbs.twimg.com/media/C4xo5cEWEAAksFH.jpg:large", "https://pbs.twimg.com/media/C4zMvyLUYAAfeZT.jpg:large"},
	TopicMessageCount: 133,
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
		t.Error("should be 'pakistan' but was ", k)
	}
}
