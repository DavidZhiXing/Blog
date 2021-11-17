class LeeterCode_24:
    def swap_pairs(self, head):
        if head is None or head.next is None:
            return head
        dummy = ListNode(0)
        dummy.next = head
        pre = dummy
        while pre.next and pre.next.next:
            first = pre.next
            second = pre.next.next
            pre.next = second
            first.next = second.next
            second.next = first
            pre = pre.next.next
        return dummy.next

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

def print_list(head):
    while head:
        print(head.val, end=' ')
        head = head.next
    print()

"""
test
"""
if __name__ == '__main__':
    l1 = ListNode(1)
    l1.next = ListNode(2)
    l1.next.next = ListNode(3)
    l1.next.next.next = ListNode(4)
    l1.next.next.next.next = ListNode(5)
    print_list(l1)
    l2 = LeeterCode_24().swap_pairs(l1)