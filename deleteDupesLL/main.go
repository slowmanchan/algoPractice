package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	low := head
	high := head.Next

	for high != nil {
		if high.Val != low.Val {
			low.Next = high
			low = low.Next
		}

		high = high.Next
	}

	low.Next = nil

	return head
}
