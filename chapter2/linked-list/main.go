package main

import (
	"fmt"
	"strings"
)

func main() {
	ll := linkedList{}
	ll.add(1)
	ll.add(2)
	ll.add(3)
	ll.add(4)
	ll.add(5)
	ll.add(6)

	fmt.Println(ll)
	fmt.Println("linked list size:", ll.size)
}

// node represents a unit of a linked list
type node struct {
	next  *node
	value int
}

func (n node) String() string {
	return fmt.Sprintf("%v", n.value)
}

type linkedList struct {
	head *node
	size int
}

func (l *linkedList) add(value int) {
	// increase counter
	l.size++

	// create new node
	newNode := node{value: value}

	if l.head == nil {
		l.head = &newNode
		return
	}

	// get tail
	current := l.head
	for ; current.next != nil; current = current.next {
	}
	current.next = &newNode
}

func (l linkedList) String() string {
	stringBuilder := strings.Builder{}

	for current := l.head; current != nil; current = current.next {
		stringBuilder.WriteString(current.String())
	}
	return stringBuilder.String()
}
