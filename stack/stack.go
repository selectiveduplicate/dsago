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
func (s *stack) Push(elements ...Types) error {
	for _, element := range elements {
		if s.top == cap(s.data) {
			return errors.New("Stack is full")
		}
		s.data[s.top] = element
		s.top++
	}
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
	aStack.Push(10, 400, 23, "Kafka", 2020, "Go")
	aStack.Pop()
	aStack.See()
}
