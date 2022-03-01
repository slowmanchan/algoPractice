package main

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{
			"normal use case",
			[]int{1, 2, 3},
			[]int{1, 2, 4},
		},
		{
			"normal use case",
			[]int{9},
			[]int{1, 0},
		},
		{
			"normal use case",
			[]int{8, 9, 9, 9},
			[]int{9, 0, 0, 0},
		},
		{
			"normal use case",
			[]int{4, 3, 2, 1},
			[]int{4, 3, 2, 2},
		},
		{
			"normal use case",
			[]int{9, 8, 9},
			[]int{9, 9, 0},
		},
		{
			"normal use case",
			[]int{9, 9, 9},
			[]int{1, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.digits)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
