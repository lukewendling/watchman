package main

import (
	"reflect"
	"testing"
)

func Test_keywordPairs_keywords(t *testing.T) {
	tests := []struct {
		name string
		kp   keywordPairs
		want []string
	}{
		{
			"keywords",
			keywordPairs{[]interface{}{"attack", 10.0}, []interface{}{"crisis", 20.0}},
			[]string{"attack", "crisis"},
		},
	}
	for _, tt := range tests {
		if got := tt.kp.keywords(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. keywordPairs.keywords() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
