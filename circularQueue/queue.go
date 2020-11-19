package main

import (
	"errors"
	"fmt"
)

// Types interface for holding various data types in the queue
type Types interface{}

// Queue defines queue data structure
type Queue struct {
	data       []Types
	head, tail int
}

func newQue(size int) *Queue {
	return &Queue{make([]Types, size), 0, 0}
}

// Enqueue enqueues elements
func (q *Queue) Enqueue(elements ...Types) (bool, error) {
	for _, element := range elements {
		if q.head == (q.tail+1)%(len(q.data)) {
			return false, errors.New("Queue full")
		}
		q.data[q.tail] = element
		q.tail = (q.tail + 1) % (len(q.data))
	}
	return true, nil
}

// Dequeue dequeues an element
func (q *Queue) Dequeue() (Types, error) {
	if q.head == q.tail {
		return 0, errors.New("Empty queue")
	}
	element := q.data[q.head]
	q.head++
	return element, nil
}

// See prints the queue
func (q *Queue) See() {
	for _, element := range q.data[q.head:] {
		fmt.Println(element)
	}
	for _, element := range q.data[0:q.tail] {
		fmt.Println(element)
	}
}

// GetHeadTail prints head and tail values
func (q *Queue) GetHeadTail() {
	fmt.Printf("Head: %v, Tail: %v\n", q.head, q.tail)
}

func main() {
	queue := newQue(5)
	queue.Enqueue("Paprika", 3.1416, 556, "Magnetic Rose", 7)
	queue.Dequeue()
	queue.Dequeue()
	queue.Enqueue(789)
	queue.Enqueue("Queue")
	queue.Dequeue()
	queue.Enqueue("Consciousness")
	queue.See()
}
