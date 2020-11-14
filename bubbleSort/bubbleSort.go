package main

import "fmt"

// BubbleSort sorts an integer array in increasing order
func BubbleSort(array []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				swapped = true
			}
		}
	}
}

func main() {
	array := []int{23, 999, 90, 455, 10, 100, 3, 5, 4, 5}
	BubbleSort(array)
	fmt.Println(array)
}
