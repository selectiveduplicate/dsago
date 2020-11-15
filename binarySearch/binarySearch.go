package main

import "fmt"

// Sort sorts an integer array for binary searching
func Sort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

// BinarySearch searches for an element in an integer array
// returns true and index no. if found, otherwise false and 0
func BinarySearch(array []int, searched int) (bool, int) {
	lowIndex := 0
	highIndex := len(array) - 1

	for lowIndex <= highIndex {
		midIndex := (lowIndex + highIndex) / 2
		if array[midIndex] == searched {
			return true, midIndex
		}
		if array[midIndex] < searched {
			lowIndex = midIndex + 1
		} else {
			highIndex = midIndex - 1
		}
	}
	return false, 0
}

func main() {
	array := []int{12, 9, 788, 401, 22, 31, 67, 8}
	Sort(array)
	fmt.Println(BinarySearch(array, 22))
}
