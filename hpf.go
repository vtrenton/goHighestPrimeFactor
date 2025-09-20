package main

import (
	"errors"
	"fmt"
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

	// Get factors up to sqrt(n) for prime calculation
	factors := relevantFactors(n)

	if len(factors) == 0 {
		return 0, errors.New("input number is prime")
	}

	pfactors := getprimes(factors)

	return pfactors[len(pfactors)-1], nil
}

// relevantFactors returns all factors of n up to sqrt(n)
// This is specifically optimized for prime factorization where larger factors aren't needed
func relevantFactors(n int) []int {
	var factors []int
	if n <= 2 {
		return factors // no factors for numbers <= 2
	}

	limit := int(math.Sqrt(float64(n)))

	// Append 2 to the list if n is even
	if n%2 == 0 {
		factors = append(factors, 2)
	}
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
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
