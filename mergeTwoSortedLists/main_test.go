package main

import (
	"reflect"
	"testing"
)

func TestMergeSortedLists(t *testing.T) {
	var tests = []struct {
		name string
		x    *ListNode
		y    *ListNode
		want *ListNode
	}{
		{
			"test 1",
			&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := mergeTwoLists(tt.x, tt.y)
			if !reflect.DeepEqual(z, tt.want) {
				t.Errorf("got %+v, want %+v", z, tt.want)
			}
		})
	}
}
