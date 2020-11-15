package main

import "fmt"

// LinearSearch searches for an element in an integer array
// returns true and index no. if found, otherwise false and 0
func LinearSearch(array []int, searched int) (bool, int) {
	for i := 0; i < len(array); i++ {
		if array[i] == searched {
			return true, i
		}
	}
	return false, 0
}

func main() {
	array := []int{12, 9, 788, 401, 22, 31, 67, 8}
	fmt.Println(LinearSearch(array, 67))
}
