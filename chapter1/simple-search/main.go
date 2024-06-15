package main

import (
	"log"
	"time"
)

const algorithmName = "simple search"

func main() {
	// MOCK SORTED LIST
	sl := sortedList(-10, 60)
	log.Printf("sorted list len: %d | sorted list: %v\n", len(sl), sl)

	// START
	log.Printf("STARTING [%s] ALGORITHM\n", algorithmName)
	start := time.Now().UTC()

	// RUNN ALGORITHM
	item := 56
	result, it := simpleSearch(sl, item)
	log.Printf("is %d present? [%v] | iterations: %d \n", item, result, it)

	// FINNISH
	elapsed := time.Since(start)
	log.Printf("FINISHED. Elapsed time: %v\n", elapsed)
}

func simpleSearch(sortedList []int, item int) (bool, int) {
	iterations := 0
	for k := range sortedList {
		iterations++
		if sortedList[k] == item {
			return true, iterations
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
