package main

import (
	"errors"
	"fmt"
	"iter"
	"math"
	"os"
	"strconv"
)

func main() {
	osargs := os.Args

	hpf, err := run(osargs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Println(hpf)
}

func run(osargs []string) (int, error) {
	if len(osargs) != 2 {
		return 0, errors.New("usage: pass in a single postive non-prime number as an Arg that you want the highest prime factorial for")
	}

	n, err := strconv.Atoi(osargs[1])
	if err != nil {
		return 0, err
	}

	if n < 0 {
		return 0, errors.New("negative numbers are not supported")
	}

	// Collect factors up to sqrt(n) for prime calculation
	var relevantFactors []int
	for factor := range factorsUpToSqrt(n) {
		relevantFactors = append(relevantFactors, factor)
	}

	if len(relevantFactors) == 0 {
		return 0, errors.New("input number is prime")
	}

	pfactors := getprimes(relevantFactors)

	return pfactors[len(pfactors)-1], nil
}

// factorsUpToSqrt returns an iterator that yields factors of n up to sqrt(n)
// This is specifically optimized for prime factorization where larger factors aren't needed
func factorsUpToSqrt(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		if n <= 2 {
			return // no factors for numbers <= 2
		}

		limit := int(math.Sqrt(float64(n)))
		for i := 2; i <= limit; i++ {
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
