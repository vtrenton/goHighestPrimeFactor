package main

import (
	"fmt"
	"math"
)

func main() {
	n := 1723343
	factors := getfactors(n)
	fmt.Println(factors)
	pfactors := getprimes(factors)
	fmt.Println(pfactors)

	//fmt.Println(pfactors[len(pfactors)-1])
}

// get all factors of n
// easiest way is to find any number that below sqrt(n)
// that x % n == 0
func getfactors(n int) []int {
	var factors []int

	limit := int(math.Sqrt(float64(n)))
	for i := 2; i < limit; i++ {
		fmt.Printf("%d\n")
		if i%n == 0 {
			factors = append(factors, i)
			fmt.Println("made it here")
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
