package main

import "fmt"

/*
You want to divide this farm evenly into square plots. You want the plots
to be as big as possible. 1680m X 640m.

How do you figure out the largest square size you can use for a plot of
land? Use the D&C strategy! D&C algorithms are recursive algorithms.
To solve a problem using D&C, there are two steps:
1. Figure out the base case. This should be the simplest possible case.
2. Divide or decrease your problem until it becomes the base case.
*/

func main() {
	farmWidth := 640
	farmLength := 1680
	fmt.Printf("the biggest square is %d meters.\n", divideFarm(farmWidth, farmLength))
}

func divideFarm(width int, length int) int {
	// base case
	if length == 2*width {
		return width
	}

	// recursive case
	// find biggest square
	n := length / width
	newWidth := length - (width * n)
	newLength := width

	return divideFarm(newWidth, newLength)
}
