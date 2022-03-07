package main

import "testing"

func TestFindTheDifference(t *testing.T) {
	var tests = []struct {
		name string
		s    string
		t    string
		want byte
	}{
		{"test 1", "abcd", "abced", 'e'},
		{"test 1", "ae", "aea", 'a'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findTheDifference(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMapFindTheDifference(t *testing.T) {
	var tests = []struct {
		name string
		s    string
		t    string
		want byte
	}{
		{"test 1", "abcd", "abced", 'e'},
		{"test 2", "ae", "aea", 'a'},
		{"test 2", "aae", "aead", 'd'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hashMapfindTheDifference(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
