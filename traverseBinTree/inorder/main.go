package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}

	n1.Right = n2
	n2.Left = n3

	inOrderTraverse(n1)
}

// func inOrderTraverse(root *TreeNode) {
// 	if root != nil {
// 		inOrderTraverse(root.Left)
// 		fmt.Println(root.Val)
// 		inOrderTraverse(root.Right)
// 	}
// }

// 1
//  \
//   2
//  /
// 3

//    f(1)  print 1
//  /      \
// f(nil)   f(2) print 2
//           /
//         f(3) print 3
//         /   \
//        f(nil) f(nil)

// using stack
func inOrderTraverse(root *TreeNode) {
	stack := []*TreeNode{}

	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = pop(&stack)
		fmt.Println(curr.Val)
		curr = curr.Right
	}
}

func pop(s *[]*TreeNode) *TreeNode {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return t
}
