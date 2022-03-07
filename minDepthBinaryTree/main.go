package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if l == 0 {
		return r + 1
	}
	if r == 0 {
		return l + 1
	}

	if l < r {
		return l + 1
	}
	return r + 1
}

// if the root is nil we are at the end of the path
// at this point we reverse up the recursion stack
// else
// l keeps recursivley calls minDepth on the left node until it
// hits no node, in which case it recurses back up the stack and
// starts resovling previous calls
// same with right

// 1
// /
// 2

// is root nil ? no (1)
// l = .. call minDepth on the left (2)
//// is root nil ? no
//// call minDepth on the left (nil)
////// is root nil ? yes return 0
//// l = 0
//// r = call minDepth on right (nil)
////// is root nil ? yes return 0
//// r = 0
//// is l == 0 ? yes
//// return 0 + 1 (1)
// l = 0
// r = call minDepth on the right (nil)
//// is root nil ? yes return 0
// r = 0
// l == 0 ? yes  
// return 0 + 1
// minPath is 1

// 1
// / \
// 2  2
// / \ 
// 3  3

// is root nil ? no (1)
// l = .. call minDepth on the left (2)
//// is root nil ? no
//// call minDepth on the left (3)
////// is root nil ? no
//////// cal minDepth on the left (nil)
//////// is root nil ? yes return 0
////// l = 0
////// r = call minDepth on right (3)
/////// is root nil ? no 
//////// r call minDepth on right (nil)
//////// is root nil ? yes return 0
////// r = 0
//// r = 0
//// is l == 0 ? yes
//// return 0 + 1 (1)
// l = 0
// r = call minDepth on the right (nil)
//// is root nil ? yes return 0
// r = 0
// l == 0 ? yes  
// return 0 + 1
// minPath is 1