__package__ = 'leetcode'

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

    def __str__(self):
        return str(self.val)

    def __repr__(self):
        return str(self.val)


def merge_two_lists(l1, l2):
    if not l1:
        return l2
    if not l2:
        return l1

    if l1.val < l2.val:
        l1.next = merge_two_lists(l1.next, l2)
        return l1
    else:
        l2.next = merge_two_lists(l1, l2.next)
        return l2



def test_merge_two_lists():
    l1 = ListNode(1)
    l1.next = ListNode(2)
    l1.next.next = ListNode(4)

    l2 = ListNode(1)
    l2.next = ListNode(3)
    l2.next.next = ListNode(4)

    l3 = merge_two_lists(l1, l2)
    assert l3.val == 1
    assert l3.next.val == 1
    assert l3.next.next.val == 2
    assert l3.next.next.next.val == 3
    assert l3.next.next.next.next.val == 4
    assert l3.next.next.next.next.next.val == 4

