package main

import (
	"reflect"
	"testing"
)

func Test_pascalsTriangle(t *testing.T) {
	tests := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{"test -1", 1, [][]int{{1}}},
		{"test -1", 3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{"test 1", 5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generate(tt.numRows)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
