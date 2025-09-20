package main

import (
	"fmt"
	"iter"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: pass in a postive number as an Arg that you want the highest prime factorial for.")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Argument was not a number")
	}

	if n < 0 {
		log.Fatal("negative numbers are not supported")
	}
	// DEBUG: test cases

	//n := 1723343
	//n := 157732
	//n := 100

	// Collect factors up to sqrt(n) for prime calculation
	var relevantFactors []int
	limit := int(math.Sqrt(float64(n)))

	for factor := range factors(n) {
		if factor > limit {
			break // stop at sqrt(n) - larger factors aren't needed for prime calculation
		}
		relevantFactors = append(relevantFactors, factor)
	}

	if len(relevantFactors) == 0 {
		fmt.Println("Input number has no factors!")
		return
	}

	pfactors := getprimes(relevantFactors)

	fmt.Println(pfactors[len(pfactors)-1])
}

// factors returns an iterator that yields factors of n one at a time
// This allows for lazy evaluation and early termination
func factors(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		if n <= 2 {
			return // no factors for numbers <= 2
		}

		for i := 2; i < n; i++ {
			if n%i == 0 {
				if !yield(i) {
					return // consumer requested early termination
				}
			}
		}
	}
}

func isprime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getprimes(factors []int) []int {
	var pfactors []int
	for _, factor := range factors {
		if isprime(factor) {
			pfactors = append(pfactors, factor)
		}
	}
	return pfactors
}
