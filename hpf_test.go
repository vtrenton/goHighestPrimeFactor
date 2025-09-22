package main

import (
	"slices"
	"testing"
)

func TestGetComposite(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		args    []string
		want    int
		wantErr bool
	}{
		{name: "valid input returns number", args: []string{"program", "100"}, want: 100, wantErr: false},
		{name: "prime number input", args: []string{"program", "97"}, want: 97, wantErr: false},
		{name: "small number input", args: []string{"program", "25"}, want: 25, wantErr: false},
		{name: "no arguments returns error", args: []string{"program"}, want: 0, wantErr: true},
		{name: "too many arguments returns error", args: []string{"program", "100", "200"}, want: 0, wantErr: true},
		{name: "non-numeric argument returns error", args: []string{"program", "abc"}, want: 0, wantErr: true},
		{name: "negative number returns error", args: []string{"program", "-5"}, want: 0, wantErr: true},
		{name: "zero is valid", args: []string{"program", "0"}, want: 0, wantErr: false},
		{name: "one is valid", args: []string{"program", "1"}, want: 1, wantErr: false},
		{name: "two is valid", args: []string{"program", "2"}, want: 2, wantErr: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := getComposite(tc.args)

			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error but got none, result: %d", got)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				} else if got != tc.want {
					t.Errorf("got %d, want %d", got, tc.want)
				}
			}
		})
	}
}

func TestRelevantFactors(t *testing.T) {
	t.Parallel()

	// Test factors up to sqrt(n)
	tests := []struct {
		name string
		in   int
		want []int
	}{
		{name: "factors of 100 up to sqrt(100) (odd numbers above 2 only)", in: 100, want: []int{2, 5}},
		{name: "factors of 36 up to sqrt(36)=6 (including 6)", in: 36, want: []int{2, 3}},
		{name: "factors of 25 up to sqrt(25)=5 (including 5)", in: 25, want: []int{5}},
		{name: "primes have no factors", in: 97, want: []int{}},
		{name: "numbers below 2 return early", in: 1, want: []int{}},
		{name: "test 2 to assure it returns early", in: 2, want: []int{}},
		{name: "test negative return early", in: -2, want: []int{}},
		{name: "small composite number 12 up to sqrt(12)≈3.46", in: 12, want: []int{2, 3}},
		{name: "perfect square 49 up to sqrt(49)=7 (including 7)", in: 49, want: []int{7}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := relevantFactors(tc.in)
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
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
		{name: "five is prime", in: 5, want: true},
		{name: "six is not prime", in: 6, want: false},
		{name: "seven is prime", in: 7, want: true},
		{name: "nine is not prime", in: 9, want: false},
		{name: "large prime (97)", in: 97, want: true},
		{name: "zero is not prime", in: 0, want: false},
		{name: "negative is not prime", in: -5, want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := isprime(tc.in); got != tc.want {
				t.Errorf("isprime(%d) = %v; want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestGetHpf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		in      []int
		want    int
		wantErr bool
	}{
		{name: "generic test", in: []int{2, 3, 4, 5, 6, 7, 8, 9}, want: 7, wantErr: false},
		{name: "unordered test (always get highest value)", in: []int{3, 12, 11, 14, 8, 7}, want: 11, wantErr: false},
		{name: "single prime", in: []int{7}, want: 7, wantErr: false},
		{name: "single composite", in: []int{8}, want: 0, wantErr: false},
		{name: "no primes", in: []int{4, 6, 8, 9}, want: 0, wantErr: false},
		{name: "empty list", in: []int{}, want: 0, wantErr: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := gethpf(tc.in)

			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error but got none, result: %d", got)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				} else if got != tc.want {
					t.Errorf("got %d, want %d", got, tc.want)
				}
			}
		})
	}
}
