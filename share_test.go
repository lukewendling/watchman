package main

import (
	_ "fmt"
	"testing"
)

func TestShareEvents(t *testing.T) {
	type args struct {
		evts events
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

func Test_stringifyGenericMap(t *testing.T) {
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
		if got := stringifyGenericMap(tt.args.evt); got != tt.want {
			t.Errorf("%q. stringifyGenericMap() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
