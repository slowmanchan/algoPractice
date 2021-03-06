package main

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want int
	}{
		{"test 1", []int{1, 2, 4, 6, 3, 3}, 3, 4},
		{"test 2", []int{1, 3, 4, 6, 3, 2}, 3, 4},
		{"test 3", []int{3, 2, 2, 3}, 3, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			if got != tt.want {
				t.Errorf("Got %d, Wanted %d", got, tt.want)
			}
		})
	}
}
