package main

import "testing"

func TestIsPrime(t *testing.T) {
	t.Run("test values under 2", func(t *testing.T) {
		got := isprime(1)
		if got { // got is set to true, but should be false
			t.Error("1 is not a prime number, but the program disagree")
		}
	})
	t.Run("validate 2 is still classified as a prime number", func(t *testing.T) {
		got := isprime(2)
		if !got { // got is set to false, but should be true
			t.Error("2 is a prime number, but the program disagrees")
		}
	})
	t.Run("validate a non-prime number", func(t *testing.T) {
		got := isprime(6)
		if got { // got is set to true, but should be false
			t.Error("6 is not a prime number, but the program disagrees")
		}
	})
	t.Run("validate a standard prime number", func(t *testing.T) {
		got := isprime(7)
		if !got { // got is set to false, but should be true
			t.Error("7 is a prime number, but the program disagrees")
		}
	})
}
