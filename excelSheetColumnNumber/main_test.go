package main

import "testing"

func Test_titleToNumber(t *testing.T) {
	tests := []struct {
		name        string
		columnTitle string
		want        int
	}{
		{"test 1", "A", 1},
		{"test 2", "AB", 28},
		{"test 2", "ZY", 701},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := titleToNumber(tt.columnTitle)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_RtoLtitleToNumber(t *testing.T) {
	tests := []struct {
		name        string
		columnTitle string
		want        int
	}{
		{"test 1", "A", 1},
		{"test 2", "AB", 28},
		{"test 2", "ZY", 701},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RtoLtitleToNumber(tt.columnTitle)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
