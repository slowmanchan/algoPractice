package main

import "testing"

func TestSqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"test 1", 2, 1},
		{"test 2", 4, 2},
		{"test 3", 8, 2},
		{"test 4", 2147395600, 46340},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mySqrt(tt.x)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
