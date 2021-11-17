const leetcodeSwapNodes = (head) => {
    if (head === null || head.next === null) {
        return head;
    }
    let prev = null;
    let curr = head;
    let next = head.next;
    while (curr !== null) {
        next = curr.next;
        curr.next = prev;
        prev = curr;
        curr = next;
    }
    return prev;
};

// test swap nodes
const testSwapNodes = () => {
    const head = new ListNode(1);
    head.next = new ListNode(2);
    head.next.next = new ListNode(3);
    head.next.next.next = new ListNode(4);
    head.next.next.next.next = new ListNode(5);
    const result = leetcodeSwapNodes(head);
    console.log(result);
};

testSwapNodes();