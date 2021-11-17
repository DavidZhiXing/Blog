package leetcode

// given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
// k is guaranteed to be a positive integer.
//
// Example 1:
//
// Given this linked list: 1->2->3->4->5
//
// For k = 2, you should return: 2->1->4->3->5
//
// For k = 3, you should return: 3->2->1->4->5
//
// Note:
//
//     If the number of nodes is less than k * n, where n is the number of nodes,
//     then return null.
//
//     You may assume k is always valid, 1 ≤ k ≤ n.
//
//     Credits:
//     Special thanks to @mithmatt for adding this problem and creating all test cases.

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		tail := cur
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		next := tail.Next
		tail.Next = nil
		cur.Next = reverse(cur.Next)
		tail.Next = next
		cur = tail
	}
	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
type ListNode struct {
	Val  int
	Next *ListNode
}

