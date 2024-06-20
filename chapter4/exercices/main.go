package main

import "fmt"

func main() {
	list := []int{1, 2, 33, 4, 5, 6}

	fmt.Printf("the sum is: %d\n", sum(list))
	fmt.Printf("the number of itens is: %d\n", itensInList(list))
	fmt.Printf("the maximum value is: %d\n", maxValue(list))

	item := 6
	fmt.Printf("is there a number %d in the list?: %t\n", item, binarySearch(list, item))
}

// exercise 4.1
func sum(list []int) int {
	// base case
	if len(list) == 1 {
		return list[0]
	}

	// recursive case
	return list[0] + sum(list[1:])
}

// exercise 4.2
func itensInList(list []int) int {
	// base case
	if len(list) == 1 {
		return 1
	}

	// recursive case
	return 1 + itensInList(list[1:])
}

// exercise 4.3
func maxValue(list []int) int {
	// base case
	if len(list) == 2 {
		if list[0] > list[1] {
			return list[0]
		}
		return list[1]
	}

	// recursive case
	subMax := maxValue(list[1:])
	if list[0] > maxValue(list[1:]) {
		return list[0]
	}
	return subMax
}

// exercise 4.4
func binarySearch(list []int, item int) bool {
	// base case
	if len(list) == 1 {
		return list[0] == item
	}

	// recursive case
	mid := len(list) / 2
	guess := list[mid]

	if guess == item {
		return true
	}
	if guess > item {
		return binarySearch(list[:mid-1], item)
	}

	return binarySearch(list[mid+1:], item)
}
