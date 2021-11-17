const { type } = require("os");

const mergeKLists = function(lists) {
    if (!lists.length) return null;
    if (lists.length === 1) return lists[0];
    let head = lists[0];
    for (let i = 1; i < lists.length; i++) {
        head = mergeTwoLists(head, lists[i]);
    }
    return head;
}

const mergeTwoLists = function(l1, l2) {
    if (!l1) return l2;
    if (!l2) return l1;
    let head = new ListNode(0);
    let cur = head;
    while (l1 && l2) {
        if (l1.val < l2.val) {
            cur.next = l1;
            l1 = l1.next;
        } else {
            cur.next = l2;
            l2 = l2.next;
        }
        cur = cur.next;
    }
    if (l1) cur.next = l1;
    if (l2) cur.next = l2;
    return head.next;
}

const ListNode = {
    val: number,
    next: ListNode
}

//test merge k lists
let l1 = {
    val: 1,
    next: {
        val: 2,
        next: {
            val: 4,
            next: null
        }
    }
}

let l2 = {
    val: 1,
    next: {
        val: 3,
        next: {
            val: 4,
            next: null
        }

    }
}

mergeKLists([l1, l2]);
