package main

import "fmt"

// InsertionSort sorts an integer array
func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		unsorted := array[i]
		for j := i - 1; j >= 0; j-- {
			if array[j] > unsorted {
				array[j+1], array[j] = array[j], unsorted
			}
		}
	}
}

func main() {
	array := []int{23, 2, 10, 50, 45, 20, 12, 33, 23, 27, 25}
	InsertionSort(array)
	fmt.Println(array)
}
