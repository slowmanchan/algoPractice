package main

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want []int
	}{
		{
			"test 1", []int{1, 2, 3, 3}, 3, []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			removeElement(tt.nums, tt.val)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("Got %d, Wanted %d", tt.nums, tt.want)
			}
		})
	}
}
