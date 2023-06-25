package ex2

import (
	"fmt"
	"testing"
)

// benchFn to be benchmarked
func benchFn(input int) {
	for i := 0; i < input; i++ {

	}
}

var testBnFnScenarios = []struct {
	input int
}{
	{input: 1},
	{input: 20},
	{input: 50},
	{input: 100},
}

func BenchmarkScenarios(b *testing.B) {
	for _, scenario := range testBnFnScenarios {
		b.Run(fmt.Sprintf("input_%d", scenario.input), func(b *testing.B) {
			for count := 0; count < b.N; count++ {
				benchFn(scenario.input)
			}
		})
	}
}
