package leetcode

// given a linked list, swap every two adjacent nodes and return its head
// you may not modify the values in the list nodes, only nodes itself may be changed
// example:
// Given 1->2->3->4, you should return the list as 2->1->4->3.

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

type ListNode struct {
	Val  int
	Next *ListNode
}