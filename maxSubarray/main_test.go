package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"test 2", []int{5, 4, -1, 7, 8}, 23},
		{"test 3", []int{1, 2}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubArray(tt.nums)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
