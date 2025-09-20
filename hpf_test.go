package main

import (
	"slices"
	"testing"
)

func TestMain(t *testing.T) {}

func TestFactors(t *testing.T) {
	t.Parallel()

	// Test complete iteration
	tests := []struct {
		name string
		in   int
		want []int
	}{
		{name: "get all factors of 100", in: 100, want: []int{2, 4, 5, 10, 20, 25, 50}},
		{name: "get all factors of 25", in: 25, want: []int{5}},
		{name: "primes have no factors", in: 97, want: []int{}},
		{name: "numbers below 2 return early", in: 1, want: []int{}},
		{name: "test 2 to assure it returns early", in: 2, want: []int{}},
		{name: "test negative return early", in: -2, want: []int{}},
		{name: "small composite number", in: 12, want: []int{2, 3, 4, 6}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var got []int
			for factor := range factors(tc.in) {
				got = append(got, factor)
			}
			if !slices.Equal(got, tc.want) {
				t.Errorf("test: %s: failed! got %v, want %v", tc.name, got, tc.want)
			}
		})
	}

	// Test early termination
	t.Run("early termination", func(t *testing.T) {
		var got []int
		for factor := range factors(100) {
			if factor >= 10 {
				break
			}
			got = append(got, factor)
		}
		want := []int{2, 4, 5}
		if !slices.Equal(got, want) {
			t.Errorf("early termination test failed! got %v, want %v", got, want)
		}
	})
}

func TestIsPrime(t *testing.T) {
	t.Parallel() // run as a goroutine to speed up tests

	tests := []struct {
		name string
		in   int
		want bool
	}{
		{name: "below 2 is not prime (1)", in: 1, want: false},
		{name: "two is prime", in: 2, want: true},
		{name: "three is prime", in: 3, want: true},
		{name: "four is not prime", in: 4, want: false},
		{name: "six is not prime", in: 6, want: false},
		{name: "seven is prime", in: 7, want: true},
		{name: "nine is not prime", in: 9, want: false},
		{name: "large prime (97)", in: 97, want: true},
		{name: "zero is not prime", in: 0, want: false},
		{name: "negative is not prime", in: -5, want: false},
	}

	for _, tc := range tests {
		tc := tc // capture for t.Parallel inside subtests
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := isprime(tc.in); got != tc.want {
				t.Errorf("isprime(%d) = %v; want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestGetPrimes(t *testing.T) {
	t.Run("")
}
