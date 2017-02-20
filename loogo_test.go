package main

import "testing"

func Test_Between(t *testing.T) {
	type args struct {
		field string
		lower string
		upper string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"between",
			args{"created", "0", "1"},
			"?filter[where][created][between][0]=0&filter[where][created][between][1]=1",
		},
	}
	for _, tt := range tests {
		if got := Between(tt.args.field, tt.args.lower, tt.args.upper); got != tt.want {
			t.Errorf("%q. Between() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_Eq(t *testing.T) {
	type args struct {
		field string
		val   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"eq",
			args{"featurizer", "image"},
			"?filter[where][featurizer]=image",
		},
	}
	for _, tt := range tests {
		if got := Eq(tt.args.field, tt.args.val); got != tt.want {
			t.Errorf("%q. Eq() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
