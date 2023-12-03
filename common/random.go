/*
 * Copyright (c) 2023 Marco Massenzio. All rights reserved.
 */

package common

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
)

// GenRandInts generates N random integers with a given distribution
func GenRandInts(n int, distribution func() int) []int {
	randomIntegers := make([]int, n)
	for i := 0; i < n; i++ {
		randomIntegers[i] = distribution()
	}
	return randomIntegers
}

// PrintRandLogs prints the log of N random integers with a uniform distribution
func PrintRandLogs(n int) {
	uniformDistribution := func() int {
		return rand.Intn(n) // Adjust the range as needed
	}

	randomIntegers := GenRandInts(n, uniformDistribution)
	slices.Sort(randomIntegers)
	// Print the log of the generated random integers
	fmt.Println("Sorted Log of Random Integers:")
	for _, randomInteger := range randomIntegers {
		fmt.Printf("Ln(%d) = %f\n", randomInteger, math.Log(float64(randomInteger)))
	}
}
