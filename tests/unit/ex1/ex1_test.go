package ex1

import (
	"fmt"
	"testing"
)

var (
	testcases = []struct {
		input    int
		expected string
	}{
		{
			input:    0,
			expected: "0",
		},
		{
			input:    1,
			expected: "1",
		},
		{
			input:    2,
			expected: "2",
		},
		{
			input:    3,
			expected: "Fizz",
		},
		{
			input:    5,
			expected: "Buzz",
		},
		{
			input:    6,
			expected: "Fizz",
		},
		{
			input:    12,
			expected: "Fizz",
		},
		{
			input:    15,
			expected: "FizzBuzz",
		},
	}
)

func TestFizzBuzz(t *testing.T) {
	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %v", i), func(t *testing.T) {
			actual := fizzBuzz(tc.input)
			if tc.expected != actual {
				t.Fatalf("Expected: %v Actual: %v", tc.expected, actual)
			}
		})
	}
}

func TestFizzBuzzBoolean(t *testing.T) {
	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %v", i), func(t *testing.T) {
			actual := fizzBuzzBooleans(tc.input)
			if tc.expected != actual {
				t.Fatalf("Expected: %v Actual: %v", tc.expected, actual)
			}
		})
	}
}
