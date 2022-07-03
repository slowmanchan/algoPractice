package main

import (
	"reflect"
	"testing"
)

func TestMergeSorted(t *testing.T) {
	type test struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}

	tests := []test{
		{
			"test 1",
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1, tt.want) {
				t.Errorf("got %d, want %d", tt.nums1, tt.want)
			}
		})
	}
}
