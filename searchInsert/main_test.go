package main

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"test 1", []int{1, 3, 5, 6}, 5, 2},
		{"test 2", []int{1, 3, 5, 6}, 2, 1},
		{"test 3", []int{1, 3, 5, 6}, 7, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInsert(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
