package main

import (
	"errors"
	"fmt"
	"iter"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	osargs := os.Args

	hpf, err := run(osargs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hpf)
}

func run(osargs []string) (int, error) {
	if len(osargs) < 2 {
		return 0, errors.New("usage: pass in a postive number as an Arg that you want the highest prime factorial for")
		os.Exit(1)
	}

	n, err := strconv.Atoi(osargs[1])
	if err != nil {
		return 0, err
	}

	if n < 0 {
		return 0, errors.New("negative numbers are not supported")
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
		return 0, errors.New("input number has no factors")
	}

	pfactors := getprimes(relevantFactors)

	return pfactors[len(pfactors)-1], nil
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
