const reserveKGroup = (head, k) => {
    if (head == null || k <= 1) return head;
    let dummy = new ListNode(0);
    dummy.next = head;
    let prev = dummy;
    let cur = head;
    let count = 0;
    while (cur) {
        let temp = cur;
        count = 0;
        while (temp && count < k) {
            temp = temp.next;
            count++;
        }
        if (count < k) break;
        let next = temp.next;
        temp.next = null;
        prev.next = reverse(cur);
        cur = next;
        prev = prev.next;
    }
    return dummy.next;
};

//test
let head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
head.next.next.next = new ListNode(4);
head.next.next.next.next = new ListNode(5);
console.log(reserveKGroup(head, 2));
console.log(reserveKGroup(head, 3));