package main

import (
	"reflect"
	"testing"
)

func TestBruteTwoSum(t *testing.T) {
	type test struct {
		name   string
		nums   []int
		target int
		want   []int
	}

	tests := []test{
		{
			name:   "test 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bruteTwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func TestTwoS(t *testing.T) {
	type test struct {
		name   string
		nums   []int
		target int
		want   []int
	}

	tests := []test{
		{
			name:   "test 1",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "test 2",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "test 3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %d, want: %d", got, tt.want)
			}
		})
	}
}
