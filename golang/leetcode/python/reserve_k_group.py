class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reserveKGroup(self, head: ListNode, k: int) -> ListNode:
        dummy = ListNode(0)
        dummy.next = head
        pre = dummy
        while pre.next:
            first = pre.next
            for i in range(k):
                if not first:
                    return dummy.next
                first = first.next
            second = first.next
            for i in range(k):
                if not first:
                    return dummy.next
                first = first.next
            pre.next = second
            first.next = second.next
            second.next = first
            pre = pre.next.next
        return dummy.next


"""
Test Case:
"""

def test_reserve_k_group():
    head = ListNode(1)
    head.next = ListNode(2)
    head.next.next = ListNode(3)
    head.next.next.next = ListNode(4)
    head.next.next.next.next = ListNode(5)
    head.next.next.next.next.next = ListNode(6)
    head.next.next.next.next.next.next = ListNode(7)
    head.next.next.next.next.next.next.next = ListNode(8)
    head.next.next.next.next.next.next.next.next = ListNode(9)
    head.next.next.next.next.next.next.next.next.next = ListNode(10)
    head.next.next.next.next.next.next.next.next.next.next = ListNode(11)
    head.next.next.next.next.next.next.next.next.next.next.next = ListNode(12)
    head.next.next.next.next.next.next.next.next.next.next.next.next = ListNode(13)
    head.next.next.next.next.next.next.next.next.next.next.next.next.next = ListNode(14)
    head.next.next.next.next.next.next.next.next.next.next.next.next.next.next = ListNode(15)
    head.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next = ListNode(16)
    head.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next.next = ListNode(17)
    Solution().reserveKGroup(head, 3)
    