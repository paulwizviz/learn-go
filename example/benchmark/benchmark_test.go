package benchmark

import (
	"fmt"
	"testing"
)

var (
	sliceData = func() []int {
		var dps []int
		for i := 0; i < 1000; i++ {
			dps = append(dps, i)
		}
		return dps
	}()

	mapData = func() map[int]int {
		m := make(map[int]int)
		for i := 0; i < 1000; i++ {
			m[i] = i
		}
		return m
	}()
)

var tests = []struct {
	input int
}{
	{input: 2},
	{input: 99},
}

func getSlice(item int) bool {
	for _, value := range sliceData {
		if value == item {
			return true
		}
	}
	return false
}

// Using slice types
func BenchmarkSlice(b *testing.B) {
	for _, test := range tests {
		b.Run(fmt.Sprintf("input_size_%d", test.input), func(b *testing.B) {
			for count := 0; count < b.N; count++ {
				getSlice(test.input)
			}
		})
	}
}

func getMap(item int) bool {
	_, ok := mapData[item]
	if !ok {
		return false
	}
	return true
}

// Switch to using map data types
func BenchmarkMap(b *testing.B) {
	for _, test := range tests {
		b.Run(fmt.Sprintf("input_size_%d", test.input), func(b *testing.B) {
			for count := 0; count < b.N; count++ {
				getMap(test.input)
			}
		})
	}
}
