package main

import (
	"errors"
	"fmt"
)

// Types interface for holding various data types in the stack
type Types interface{}

// Stack defines a stack data structure
type stack struct {
	data []Types
	top  int
}

// returns pointer to a new stack with requested size
func newStack(size int) *stack {
	return &stack{make([]Types, size), 0}
}

// pushes an element into stack
func (s *stack) Push(element Types) error {
	if s.top == cap(s.data) {
		return errors.New("Stack is full")
	}
	s.data[s.top] = element
	s.top++
	return nil
}

// pops an element
func (s *stack) Pop() (Types, error) {
	if s.top == 0 {
		return 0, errors.New("Empty stack")
	}
	element := s.data[s.top-1]
	s.top--
	return element, nil
}

// prints the stack
func (s *stack) See() {
	for _, elements := range s.data[:s.top] {
		fmt.Println(elements)
	}
}

func main() {
	aStack := newStack(5)
	aStack.Push(200)
	aStack.Push(400.23)
	aStack.Push(350)
	aStack.Push("Kafka")
	aStack.Push(5324)
	aStack.Pop()
	aStack.See()
}
