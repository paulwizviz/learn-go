package main

import "testing"

func FuzzDivideByZero(f *testing.F) {
	testcases := []int{
		0, 1, 3, 4,
	}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, n int) {
		actual := DivideByZero(n)
		expected := float64(n / (n - 2))
		if expected != actual {
			t.Errorf("Expected: %v Actual: %v", expected, actual)
		}
	})

}
