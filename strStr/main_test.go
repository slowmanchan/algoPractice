package main

import "testing"

func Test_strStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{"test 1", "hello", "ll", 2},
		{"test 2", "aaaaa", "bba", -1},
		{"test 3", "", "", 0},
		{"test 4", "aaa", "aaaa", -1},
		{"test 4", "mississippi", "mississippi", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strStr(tt.haystack, tt.needle)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
