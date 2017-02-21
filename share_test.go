package main

import (
	_ "fmt"
	"testing"
)

func TestShareEvents(t *testing.T) {
	type args struct {
		evts []event
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		ShareEvents(tt.args.evts)
	}
}

func Test_stringifyEvent(t *testing.T) {
	type args struct {
		evt map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := stringifyEvent(tt.args.evt); got != tt.want {
			t.Errorf("%q. stringifyEvent() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
