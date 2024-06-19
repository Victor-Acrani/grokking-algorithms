package main

import (
	"fmt"
	"strings"
)

func main() {
	var s stack
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(s)

	s.pop()
	fmt.Println(s)
	s.pop()
	fmt.Println(s)
}

type node struct {
	value int
	next  *node
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.value)
}

type stack struct {
	size int
	top  *node
}

func (s *stack) push(value int) {
	s.size++

	// create new node
	newNode := node{value: value}

	// check if top is nil
	if s.top == nil {
		s.top = &newNode
		return
	}

	// set new stack top
	oldTop := s.top
	newNode.next = oldTop
	s.top = &newNode
}

func (s *stack) pop(){
	s.size--

	// check if top is nil
	if s.top == nil{
		return
	}

	// set new top
	s.top = s.top.next
}

func (s stack) String() string {
	stringBuilder := strings.Builder{}

	for top := s.top; top != nil; top = top.next{
		stringBuilder.WriteString(fmt.Sprintf("%d", top.value))
	}

	return stringBuilder.String()
}
