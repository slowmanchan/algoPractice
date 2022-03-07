package main

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		node *TreeNode
		want int
	}{
		{
			"test 1",
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDepth(tt.node)
			if got != tt.want {
				t.Errorf("Got %d, Wanted %d", got, tt.want)
			}
		})
	}
}
