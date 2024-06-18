package main

import (
	"log"
	"time"
)

const algorithmName = "factorial"

func main() {
	// START
	log.Printf("STARTING [%s] ALGORITHM\n", algorithmName)
	start := time.Now().UTC()

	// RUN ALGORITHM
	x := 10
	reasult := factorial(x)
	log.Printf("%d! = %d \n", x, reasult)

	// FINNISH
	elapsed := time.Since(start)
	log.Printf("FINISHED. Elapsed time: %v\n", elapsed)
}

func factorial(x int) int {
	// base case
	if x == 0 {
		return 1
	}

	// recursive case
	return x * factorial(x-1)
}
