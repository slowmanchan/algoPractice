package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	type test struct {
		name string
		nums []int
		want int
	}

	tests := []test{
		{"test 1", []int{0, 0, 1, 1}, 2},
		{"test 1", []int{1, 1, 2}, 2},
		{"test 1", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDupes(tt.nums)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
