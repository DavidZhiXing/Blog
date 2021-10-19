package main

import "fmt"

type Node Struct {
    prev *Node
    next *Node
    key interface{}
}

type List struct {
    head *Node
    tail *Node
}

fun (L *List) Insert(key interface{}) {
    list := &Node {
        next: L.head,
        key: key,
    }
    if L.head != nil {
        L.head.prev = list
    }
    L.head = list

    l := L.head
    for l.next != nil {
        l = l.next
    }
    L.tail = l
}

fun Display(list *Node) {
    for list != nil {
        fmt.Printf("%v <-", list.key)
        list = list.prev
    }
    fmt.Println()
}

func (l *List) Reverse() {
    curr := l.head
    var prev *Node
    l.tail = l.head

    for curr != nil {
        next := curr.next
        curr.next = prev
        prev = curr
        curr = next
    }
    l.head = prev
    Display(l.head)
}

func main() {
	link := List{}
	link.Insert(5)
	link.Insert(9)
	link.Insert(13)
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)
	
	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.key)
	fmt.Printf("Tail: %v\n", link.tail.key)
	link.Display()
	fmt.Println("\n==============================\n")
	fmt.Printf("head: %v\n", link.head.key)
	fmt.Printf("tail: %v\n", link.tail.key)
	link.Reverse()
	fmt.Println("\n==============================\n")
}