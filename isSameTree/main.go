package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil || p != nil && q == nil {
		return false
	}

	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

// recursive with short circuit evaluation
// check if both sides are nil. This means that we either have no
// trees OR we are at the end which both means the trees are the same
// if either side is nil then the trees are not balanced
// we check the values of each, if they are not then we dont recurse
// because of short circuit eval. Go we'll see that the first bool is false
// and bail.
// if they are the same lets run the function again on the left side
// we go all the way down the left recursively.
// again, with short circuit, if they all pass, we go all the way down the right
