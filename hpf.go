package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("vim-go")
}

// get all factors of n
// easiest way is to find any number that below sqrt(n)
// that x % n == 0
func getfactors(n int) []int {
	var factors []int

	limit := math.Sqrt(n)
	for i := 2; i < limit; i++ {
		if i%n == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func isprime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getprimes(factors []int) []int {
	for i := range factors {

	}
}
