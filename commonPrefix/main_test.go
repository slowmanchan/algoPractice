package main

import "testing"

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		name string
		x    []string
		want string
	}{
		{"test 1", []string{"flower", "flow", "flight"}, "fl"},
		{"test 2", []string{"ab", "a"}, "a"},
		{"test 3", []string{"reflower", "flow", "flight"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lenOfLongword := longestCommonPrefix(tt.x)
			if lenOfLongword != tt.want {
				t.Errorf("got %v, want %v", lenOfLongword, tt.want)
			}
		})
	}
}
