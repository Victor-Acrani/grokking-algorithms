package main

import (
	"log"
	"time"
)

const algorithmName = "binary search"

func main() {
	// MOCK SORTED LIST
	sl := sortedList(-10, 60)
	log.Printf("sorted list len: %d | sorted list: %v\n", len(sl), sl)

	// START
	log.Printf("STARTING [%s] ALGORITHM\n", algorithmName)
	start := time.Now().UTC()

	// RUN ALGORITHM
	item := 56
	result, it := binarySearch(sl, item)
	log.Printf("is %d present? [%v] | iterations: %d \n", item, result, it)

	// FINNISH
	elapsed := time.Since(start)
	log.Printf("FINISHED. Elapsed time: %v\n", elapsed)
}

func binarySearch(sortedList []int, item int) (bool, int) {
	iterations := 0
	minIndex := 0
	maxIndex := len(sortedList) - 1

	for minIndex <= maxIndex {
		iterations++
		mid := (maxIndex + minIndex) / 2
		guess := sortedList[mid]  

		// check if mean is the desired value
		if guess == item {
			return true, iterations
		}
		// check guess and define new range
		if guess > item {
			maxIndex = mid - 1
		} else {
			minIndex = mid + 1
		}
	}
	return false, iterations
}

func sortedList(minValue int, maxValue int) []int {
	// define list size
	listsize := maxValue - minValue + 1

	// set list
	currentValue := minValue

	list := make([]int, listsize)
	for i := 0; i < listsize; i++ {
		list[i] = currentValue
		currentValue++
	}
	return list
}
