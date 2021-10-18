
node addNode(node head, int value) {
    node temp, p;
    temp = createNode();
    temp->data = value;
    if(head == NULL) {
        head = temp;
    } else {
        p = head;
        while (p->next != NULL) {
            p=p->next;
        }
        p->next = temp;
    }
    return head;
}

bool Contains(node head, int value) {
    node p;
    p = head;
    while(p != null && p->data != value){
        p = p->next;
    }
    if(p==NULL) return false;
    return true;
}

node Prepend(node head, node tail, int value) {
    node temp;
    temp = createNode(value);
    temp.next = head;
    head = temp;
    if(tail != NULL) tail = temp;
    return head;
}

bool Remove(node head, node tail, int value) {
    node p, temp;
    if(head == NULL) return false;
    p = head;
    if(p->data == value) {
        if(head == tail){
            head = NULL;
            tail = NULL;
        } else {
            head = head.next
            return true
        }
    }
    while(p->next != null and p->data != value) {
        p = p->next
    }
    if(p -> next != null) {
        if(p->next == tail) {
            tail = p;
        }
        p->next = p->next->next;
        return true
    }
    return false;
}