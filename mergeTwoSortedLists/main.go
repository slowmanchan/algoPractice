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
	// setup a dummy pointer to keep track of the head
	// if not, when you traverse, you lose track of where the front is
	// keep track (point to) the previous node as well (just copy the dummyhead)
	dummyHead := new(ListNode)
	prev := dummyHead

	// keep going until both lists are exhausted of nodes to check
	for l1 != nil && l2 != nil {
		// check the smaller value
		if l1.Val < l2.Val {
			// if its l1, set the next node to l1 (we are sorting small to high)
			// we dont want to use dummy because thats the head we need to return
			// thats why we have a prev node pointer so we set that
			prev.Next = l1
			// move the l1 node along since we've check and "used" the value
			// if we dont move, we will keep checking the same value
			l1 = l1.Next
		} else {
			// same as above
			prev.Next = l2
			l2 = l2.Next
		}
		// we ALWAYS move up the prev pointer. If we dont
		// we will be stuck AND override the nodes everytime
		prev = prev.Next
	}

	// we might have leftover nodes. Those values might
	// have been larger then all the other nodes so we just
	// merge it in by checking if they are null.
	// the list node would be null because we kept moving the pointer
	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}

	// make sure the return Next node in the dummy head
	// because when we initialized it, it just had a dummy value (nil)
	// we were actually just setting the next node over and over.
	return dummyHead.Next
}
