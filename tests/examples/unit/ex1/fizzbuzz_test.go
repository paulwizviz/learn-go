package ex1

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testcases := []struct {
		input       uint
		expected    string
		description string
	}{
		{
			input:       0,
			expected:    "0",
			description: "Expected: 0 Actual: ",
		},
		{
			input:       1,
			expected:    "1",
			description: "Expected: 1 Actual: ",
		},
		{
			input:       2,
			expected:    "2",
			description: "Expected: 2 Actual: ",
		},
		{
			input:       3,
			expected:    "Fizz",
			description: "Expected: Fizz Actual: ",
		},
		{
			input:       5,
			expected:    "Buzz",
			description: "Expected: Buzz Actual: ",
		},
		{
			input:       6,
			expected:    "Fizz",
			description: "Expected: Fizz Actual: ",
		},
		{
			input:       12,
			expected:    "Fizz",
			description: "Expected: Fizz Actual: ",
		},
		{
			input:       15,
			expected:    "Fizz Buzz",
			description: "Expected: Fizz Buzz Actual: ",
		},
	}
	for i, tc := range testcases {
		actual := fizbuz(tc.input)
		if tc.expected != actual {
			t.Fatalf("Case: %v - %v %v", i, tc.description, actual)
		}
	}
}
