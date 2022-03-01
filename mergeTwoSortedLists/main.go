package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// if list 1 is nil, its exhausted check the other
	if l1 == nil {
		return l2
	}
	// if list 2 is nil, its exhausted check the other
	if l2 == nil {
		return l1
	}
	// move to the next list value (whichevers lower) and replay the higher list
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

// set the next list and return the next non exhausted one, the function will always return
// a non empty list because of the two checks up top. We are pretty much stuck in both recursion callss
// until BOTH lists are empty.
// these returned lists are being built up and added to for the result

// Example:
// Check both lists (make sure they are not empty)
// Check the diff between both values
// L1 is less or equal so move to the next L1 list (l1.next)
// Doing this will order the recursive calls and will resolve with
// each list being added to the next

// No recursion
func mergeTwoListsNoRecurse(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = new(ListNode)
	var p = dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}
	return dummy.Next
}
