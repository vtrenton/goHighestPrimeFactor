package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var err error

	// get the composite from the Args
	var composite int
	composite, err = getComposite(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// Get factors up to sqrt(n) for prime calculation
	var factors []int
	factors, err = relevantFactors(composite)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// Get the highest factor that is prime and return it.
	var hpf int
	hpf, err = gethpf(factors)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Println(hpf)
}

func getComposite(osargs []string) (int, error) {
	// validate there are the correct amount or arguments
	// else document usage to user
	if len(osargs) != 2 {
		return 0, errors.New("usage: pass in a single postive non-prime number as an Arg that you want the highest prime factorial for")
	}
	// try to convert osArg to int
	n, err := strconv.Atoi(osargs[1])
	if err != nil {
		return 0, err
	}
	// validate n is not a negative number
	if n < 0 {
		return 0, errors.New("negative numbers are not supported")
	}

	return n, err
}

// relevantFactors returns all factors of n up to sqrt(n)
// This is specifically optimized for prime factorization where larger factors aren't needed
func relevantFactors(n int) ([]int, error) {
	var factors []int
	if n < 2 {
		return []int{}, errors.New("numbers under two have no factors")
	}

	limit := int(math.Sqrt(float64(n)))

	// Append 2 to the list if n is even (but not if n is 2 itself)
	if n%2 == 0 && n > 2 {
		factors = append(factors, 2)
	}
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors, nil
}

// get the last and highest value in the hpf
func gethpf(factors []int) (int, error) {
	// handle case of empty list (likely means composite was prime)
	if len(factors) == 0 {
		return 0, errors.New("input number is prime")
	}

	var hpf int
	for _, factor := range factors {
		if isprime(factor) && factor > hpf { // handle a bug if the list is unsorted for some reason
			hpf = factor
		}
	}
	return hpf, nil
}

// helper func for gethpf
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
