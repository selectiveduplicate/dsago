package priorityQueue

import "github.com/selectiveduplicate/dsago/heap"

func extractMax(maxHeap *[]int, heapSize int) int {
	// since max heap the first node is the maximum
	// priority node
	maxItem := (*maxHeap)[1]
	// bring the last element in root
	// and decrease heapSize
	(*maxHeap)[1] = (*maxHeap)[heapSize]
	heapSize--
	*maxHeap = (*maxHeap)[:heapSize+1]
	heap.MaxHeapify(maxHeap, heapSize, 1)

	return maxItem
}

//inserts a new element
//the new element doesn't have any associated priority
//being max heap, after inserting, it is sent to the
//appropriate position within the tree
func insertNode(maxHeap *[]int, heapSize int, new int) int {
	// increase heapSize
	heapSize++
	// insert the new item
	*maxHeap = append(*maxHeap, new)
	i := heapSize
	// to maintain max heap property
	// start checking up the tree
	for i > 1 && heap.ParentIndex(i) < (*maxHeap)[i] {
		parent := heap.ParentIndex(i)
		(*maxHeap)[i], (*maxHeap)[parent] = (*maxHeap)[parent], (*maxHeap)[i]
		i = parent
	}

	return heapSize
}
