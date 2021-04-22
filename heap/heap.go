package heap

func leftNodeIndex(i int) int {
	return 2 * i
}

func rightNodeIndex(i int) int {
	return 2*i + 1
}

func parentIndex(i int) int {
	return i / 2
}

// converts a heap into a max heap
func maxHeapify(heap *[]int, heapSize int, i int) {
	var largestNodeIndex int

	left := leftNodeIndex(i)
	right := rightNodeIndex(i)

	// left node is larger than given node
	if left <= heapSize && (*heap)[left] > (*heap)[i] {
		largestNodeIndex = left
	} else {
		largestNodeIndex = i
	}

	// right node is the largest
	if right <= heapSize && (*heap)[right] > (*heap)[largestNodeIndex] {
		largestNodeIndex = right
	}

	// if the subtree is already a max heap then we're done
	if largestNodeIndex != i {
		(*heap)[i], (*heap)[largestNodeIndex] = (*heap)[largestNodeIndex], (*heap)[i]
		maxHeapify(heap, heapSize, largestNodeIndex)
	}
}

// builds a max heap from a given heap
func buildMaxHeap(heap *[]int, heapSize int) {
	for i := heapSize / 2; i >= 1; i-- {
		maxHeapify(heap, heapSize, i)
	}
}
