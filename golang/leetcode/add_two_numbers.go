// you are given two non-empty linked lists representing two non-negative integers.
// the digits are stored in reverse order and each of their nodes contain a single digit.
// add the two numbers and return it as a linked list.

// you may assume the two numbers do not contain any leading zero, except the number 0 itself.
// example:
// input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// output: 7 -> 0 -> 8
// explanation: 342 + 465 = 807.

package main

import (
	"structures"
)

// ListNode is a list node
type ListNode = structures.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		carry int
		head  *ListNode
		tail  *ListNode
	)

	for l1 != nil || l2 != nil || carry != 0 {
		var (
			sum int
			node *ListNode
		)

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry

		carry = sum / 10
		sum = sum % 10

		node = &ListNode{Val: sum}

		if head == nil {
			head = node
		} else {
			tail.Next = node
		}

		tail = node
	}

	return head

