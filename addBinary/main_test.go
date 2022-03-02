package main

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{"test 1", "11", "0", "11"},
		{"test 2", "11", "1", "100"},
		{"test 3", "10", "1", "11"},
		{"test 4", "1010", "1011", "10101"},
		{"test 5", "1", "111", "1000"},
		{"test 6", "100", "110010", "110110"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addBinary(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
