package main

import "testing"

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		name string
		x    int
		want bool
	}{
		{"test 1", 121, true},
		{"test 2", -121, false},
		{"test 3", 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lenOfLongword := isPalindrome(tt.x)
			if lenOfLongword != tt.want {
				t.Errorf("got %v, want %v", lenOfLongword, tt.want)
			}
		})
	}
}
