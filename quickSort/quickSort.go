package main

import "fmt"

// QuickSort sorts an integer array
func QuickSort(array []int, low, high int) {
	if low >= high {
		return
	}
	pivotIndex := partition(array, low, high)
	QuickSort(array, low, pivotIndex-1)
	QuickSort(array, pivotIndex+1, high)
}

func partition(array []int, low, high int) int {
	var i, j int
	// last element as pivot
	pivot := array[high]
	for i, j = low-1, low; j < high; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[high] = array[high], array[i+1]
	return i + 1
}

func main() {
	array := []int{23, 999, 90, 455, 10, 100, 3, 5, 4, 5}
	QuickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
