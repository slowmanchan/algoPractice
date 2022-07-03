package main

import "testing"

func TestLongest(t *testing.T) {
	type test struct {
		name string
		s    string
		want int
	}

	tests := []test{
		{
			"test 1", "dvdf", 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
