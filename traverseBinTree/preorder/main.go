package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	// make a stack of node pointers
	stack := []*TreeNode{}
	output := []int{}

	// guard on nils
	if root == nil {
		return output
	}

	// throw on the first node to the stack
	stack = append(stack, root)

	// quit when we have no more nodes to process
	for len(stack) > 0 {
		// pop off a node to process
		node := pop(&stack)

		// since its preorder we are going
		// to populate the value the first time its seen
		output = append(output, node.Val)

		// must go right for pre order first
		// put this node to be processed
		// the goal is to put all the nodes on the stack
		// in the order we want to visit them
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		// put on the left
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	// return the output
	return output
}

func pop(s *[]*TreeNode) *TreeNode {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}
