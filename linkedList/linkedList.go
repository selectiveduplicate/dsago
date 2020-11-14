package main

import (
	"errors"
	"fmt"
)

// Node of a linked list
type Node struct {
	data     int
	nextNode *Node
}

// LinkedList class
type LinkedList struct {
	head *Node
}

// AddFirst adds a node at the start of the list
func (linkedList *LinkedList) AddFirst(data int) {
	node := Node{}
	node.data = data
	if node.nextNode == nil {
		node.nextNode = linkedList.head
	}
	linkedList.head = &node
}

// AddLast adds a node at the end of the list
func (linkedList *LinkedList) AddLast(data int) {
	node := Node{}
	node.data = data
	current := linkedList.head
	for current.nextNode != nil {
		current = current.nextNode
	}
	current.nextNode = &node
}

// RemoveByVal removes an item based on the data
func (linkedList *LinkedList) RemoveByVal(data int) (bool, error) {
	if linkedList.head == nil {
		return false, errors.New("Empty list")
	}
	if linkedList.head.data == data {
		linkedList.head = linkedList.head.nextNode
		return true, nil
	}
	current := linkedList.head
	for current.nextNode != nil {
		if current.nextNode.data == data {
			current.nextNode = current.nextNode.nextNode
			return true, nil
		}
		current = current.nextNode
	}
	return false, errors.New("Not found")
}

// RemoveByIndex removes an item based on the index given
func (linkedList *LinkedList) RemoveByIndex(index int) (bool, error) {
	if linkedList.head == nil {
		return false, errors.New("Empty list")
	}
	if index < 0 {
		return false, errors.New("Index cannot be negative")
	}
	if index == 0 {
		linkedList.head = linkedList.head.nextNode
		return true, nil
	}
	current := linkedList.head
	for i := 1; i < index; i++ {
		// don't bother if we've reached the end
		if current.nextNode.nextNode == nil {
			return false, nil
		}
		current = current.nextNode
	}
	current.nextNode = current.nextNode.nextNode
	return true, nil
}

// GetFirst returns the first node's data
func (linkedList *LinkedList) GetFirst() (int, error) {
	if linkedList.head == nil {
		return 0, errors.New("Empty list")
	}
	return linkedList.head.data, nil
}

// GetLast returns the last node's data
func (linkedList *LinkedList) GetLast() (int, error) {
	if linkedList.head == nil {
		return 0, errors.New("Empty list")
	}
	current := linkedList.head
	for current.nextNode != nil {
		current = current.nextNode
	}
	return current.data, nil
}

// SearchData searches for a data in the list
func (linkedList LinkedList) SearchData(searched int) (bool, error) {
	if linkedList.head == nil {
		return false, errors.New("Empty list")
	}
	current := linkedList.head
	for current.nextNode != nil {
		if current.data == searched {
			return true, nil
		}
		current = current.nextNode
	}
	return false, nil
}

// GetSize returns the size of the list
func (linkedList *LinkedList) GetSize() int {
	var size int
	current := linkedList.head
	for current.nextNode != nil {
		current = current.nextNode
		size++
	}
	return size + 1
}

// GetData returns a slice containing list data
func (linkedList *LinkedList) GetData() []int {
	var items []int
	current := linkedList.head
	for current != nil {
		items = append(items, current.data)
		current = current.nextNode
	}
	return items
}

func main() {
	linkedList := LinkedList{}
	linkedList.AddFirst(10)
	linkedList.AddFirst(250)
	linkedList.AddLast(400)
	linkedList.AddLast(778)
	linkedList.AddFirst(900)
	fmt.Println(linkedList.GetData())
}
