package main

import (
	"fmt"
	"log"
)

const algorithmName = "quicksort"

func main() {
	// MOCK UNSORTED LIST
	ul := unsortedList()

	// START
	log.Printf("STARTING [%s] ALGORITHM\n", algorithmName)
	fmt.Printf("len: %d | unsorted list: %v\n", len(ul), ul)
	sl := quickSort(ul)
	fmt.Printf("len: %d | sorted list: %v\n", len(sl), sl)
}

func unsortedList() []int {
	return []int{4, 7, 3, 8, 9, 33, 6, 5, 2, 67, 54, 23, 55}
}

func quickSort(ul []int) []int {
	// base case
	if len(ul) < 2 {
		return ul
	}

	// recursive case
	// get pivot
	pivot := len(ul) / 2

	// partition
	var smaller []int
	var bigger []int

	for k := range ul {
		if k == pivot {
			continue
		}

		if ul[k] > ul[pivot] {
			bigger = append(bigger, ul[k])
		}
		if ul[k] < ul[pivot] {
			smaller = append(smaller, ul[k])
		}
	}

	return append(append(quickSort(smaller), ul[pivot]), quickSort(bigger)...)
}
