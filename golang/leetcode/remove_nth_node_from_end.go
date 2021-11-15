package leetcode

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
//
// Given a linked list, remove the n-th node from the end of list and return its head.
//
// Example:
//
// Given linked list: 1->2->3->4->5, and n = 2.
//
// After removing the second node from the end, the linked list becomes 1->2->3->5.
// Note:
//
// Given n will always be valid.
//
// Follow up:
//
// Could you do this in one pass?


type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	var dummy = &ListNode{
		Val:  0,
		Next: head,
	}

	var fast = dummy
	var slow = dummy

	// move fast n steps ahead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// move fast and slow at the same time
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// remove slow.Next
	slow.Next = slow.Next.Next

	
	return dummy.Next
}