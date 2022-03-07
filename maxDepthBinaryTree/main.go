package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	return 1 + max(l, r)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// find the max depth of a binary tree
// check if we have a binary tree, return 0 if we do not
// keep track of the walks on left
// keep track of the walks on right
// recursively run on each side to cover both left and right for each node
// return 1 + the greater of left or right

// EG
//   1
//  / \
// 2  3

// root is not nil
// call maxDepth on the left node storing it in l. node left.value = 2
// root is not nil
// call maxDepth on the left node, node is nil
// root is nil return 0 we are at the end, start reversing and resolving
// l = 0 in the previous stack
// call maxDepth on right, node is nil
// return 0
// r = 0 in the previous call
// max = 0
// return 1
// l = 1 in the previous previous stack
// call maxDepth on the right node storing it in l. node left.value = 2
// root is not nil
// call maxDepth on the left node, node is nil
// root is nil return 0
// l = 0 in the previous stack
// call maxDepth on right, node is nil
// return 0
// r = 0 in the previous stack
// max = 0
// return 1
// l = 1 in the previous previous stack

// max of l and r is 1
// add 1 for the final result
