package main

import (
	_ "fmt"
	"testing"
)

func TestShareEvent(t *testing.T) {
	type args struct {
		evt *event
	}

	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}

	for _, tt := range tests {
		ShareEvent(tt.args.evt)
	}
}
