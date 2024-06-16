package main

import (
	"log"
	"time"
)

const algorithmName = "selection sort"

func main() {
	// MOCK LIST
	l := list()
	log.Printf("list: %v\n", l)

	// START
	log.Printf("STARTING [%s] ALGORITHM\n", algorithmName)
	start := time.Now().UTC()

	// RUN ALGORITHM
	result, it := selectionSort(l)
	log.Printf("sortedlist: %v | iterations: %d \n", result, it)

	// FINNISH
	elapsed := time.Since(start)
	log.Printf("FINISHED. Elapsed time: %v\n", elapsed)
}

func selectionSort(list []int) ([]int, int) {
	iterations := 0
	sortedList := make([]int, len(list))
	var sortedIndex int
	var value int

	for len(list) != 0 {
		value, list = findSmallest(list)
		sortedList[sortedIndex] = value
		sortedIndex++
		iterations++
	}

	return sortedList, iterations
}

func findSmallest(list []int) (int, []int) {
	smallest := list[0]
	var smallestIndex int

	// find smallest element
	for i := 1; i < len(list); i++ {
		if list[i] < smallest {
			smallest = list[i]
			smallestIndex = i
		}
	}

	// pop element from slice
	newList := append(list[:smallestIndex], list[smallestIndex+1:]...)

	return smallest, newList
}

func list() []int {
	return []int{156, 141, 35, 94, 88, 61, 111}
}
