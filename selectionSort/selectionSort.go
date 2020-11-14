package main

import "fmt"

// SelectionSort sorts an integer slice in increasing order
func SelectionSort(array []int) {
	for i := range array {
		minimum := array[i]
		minimumIndex := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < minimum {
				minimum = array[j]
				minimumIndex = j
			}
		}
		array[minimumIndex], array[i] = array[i], array[minimumIndex]
	}
}

func main() {
	array := []int{23, 999, 90, 455, 10, 100, 3, 5, 4, 5}
	SelectionSort(array)
	fmt.Println(array)
}
