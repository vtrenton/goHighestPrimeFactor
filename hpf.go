package main

import (
	"fmt"
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

	factors := getfactors(n)
	if len(factors) == 0 {
		fmt.Println("Input number has no factors!")
		return
	}

	pfactors := getprimes(factors)

	fmt.Println(pfactors[len(pfactors)-1])
}

func getfactors(n int) []int {
	var factors []int
	if n <= 2 { //intentionally not supporting negative numbers for now
		return factors //early return for inputs between -2 and 2
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 2; i < limit; i++ {
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
	for factor := range factors {
		if isprime(factor) {
			pfactors = append(pfactors, factor)
		}
	}
	return pfactors
}
