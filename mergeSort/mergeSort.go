package main

import "fmt"

// Merge merges two sorted subarrays
func Merge(array []int, low, high, middle int) {
	var leftArray, rightArray []int
	for i := low; i <= middle; i++ {
		leftArray = append(leftArray, array[i])
	}
	for i := middle + 1; i <= high; i++ {
		rightArray = append(rightArray, array[i])
	}
	a := low
	l, r := 0, 0
	for l < len(leftArray) && r < len(rightArray) {
		if leftArray[l] <= rightArray[r] {
			array[a] = leftArray[l]
			l++
		} else {
			array[a] = rightArray[r]
			r++
		}
		a++
	}
	// in case some elements are left
	if l < len(leftArray) {
		for _, element := range leftArray[l:] {
			array[a] = element
			a++
		}
	} else {
		for _, elements := range rightArray[r:] {
			array[a] = elements
			a++
		}
	}
}

// MergeSort recursively calls itself to sort an integer array
func MergeSort(array []int, low, high int) {
	var middle int
	if low < high {
		middle = low + (high-low)/2
		MergeSort(array, low, middle)
		MergeSort(array, middle+1, high)
		Merge(array, low, high, middle)
	}
}

func main() {
	array := []int{100, 50, 12, 500, 43, 78, 10, 20, 5, 200, 5}
	MergeSort(array, 0, len(array)-1)
	fmt.Println(array)
}
