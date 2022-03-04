package main

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{3, 2, 3}, 3},
		{"test 2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"test 2", []int{1, 2, 3, 4, 4}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.nums)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

// most frequent number in array
// run through array linearly
// keep track of number of times each number is visited (a map? or dict?) key = num, v = freq
// go through the map and keep track of the most frequent (v) and num (key)
// return the largest
