package main

import "testing"

// This fuzz is intended to determine if the function
// DivideByZero will panic.
//
// To execute the fuzz test, run the command
//
//	go test -fuzz Fuzz
func FuzzDivideByZero(f *testing.F) {
	// This sets up the seed values for the fuzzy
	// operator to generate values based on this
	// pattern.
	testcases := []int{
		0, 1, 3, 4,
	}
	for _, tc := range testcases {
		f.Add(tc)
	}
	// Fuzz test called by injecting value
	// n until panic is called.
	// If a value causes the function to panic
	// it will generate a folder named "testdata/fuzz/FuzzDivideByZero"
	// and store the value in the testdata
	f.Fuzz(func(t *testing.T, n int) {
		DivideByZero(n)
	})

}
