package bigo

import (
	"fmt"
	"testing"
)

var (
	mapData = func() map[int]int {
		m := make(map[int]int)
		for i := 0; i < 10_000; i++ {
			m[i] = i
		}
		return m
	}()

	sliceData = func() []int {
		var slice []int
		for value := 0; value < 10_000; value++ {
			slice = append(slice, value)
		}
		return slice
	}()
)

var gettests = []struct {
	input int
}{
	{input: 1},
	{input: 20},
	{input: 50},
	{input: 100},
	{input: 900},
	{input: 1050},
	{input: 5000},
	{input: 9999},
}

func mapItemExist(item int) bool {
	_, ok := mapData[item]
	if !ok {
		return false
	}
	return true
}

// O(1)
func BenchmarkO1(b *testing.B) {
	for _, test := range gettests {
		b.Run(fmt.Sprintf("input_%d", test.input), func(b *testing.B) {
			for count := 0; count < b.N; count++ {
				mapItemExist(test.input)
			}
		})
	}
}

// O(N)
func sliceItemExist(item int) bool {
	for _, value := range sliceData {
		if value == item {
			return true
		}
	}
	return false
}

func BenchmarkON(b *testing.B) {
	for _, test := range gettests {
		b.Run(fmt.Sprintf("input_%d", test.input), func(b *testing.B) {
			for count := 0; count < b.N; count++ {
				sliceItemExist(test.input)
			}
		})
	}
}

// O(N^2)
var dupTest = []struct {
	d1 []int
	d2 []int
	tc int
}{
	{
		d1: []int{1, 2},
		d2: []int{1, 2},
		tc: 1,
	},
	{
		d1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		d2: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		tc: 2,
	},
	{
		d1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		d2: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		tc: 3,
	},
}

func isDupSlice(d1, d2 []int) bool {
	for oIndex, outer := range d1 {
		for iIndex, inner := range d2 {
			if oIndex == iIndex {
				if outer != inner {
					return false
				}
			}
		}
	}
	return true
}

func BenchmarkN2(b *testing.B) {
	for _, test := range dupTest {
		b.Run(fmt.Sprintf("input_%d", test.tc), func(b *testing.B) {
			for count := 0; count < b.N; count++ {
				isDupSlice(test.d1, test.d2)
			}
		})
	}
}

// O(2^N)

var counttests = []struct {
	input int
}{
	{
		input: 1,
	},
	{
		input: 50,
	},
	{
		input: 100,
	},
	{
		input: 1_000,
	},
}

func countToZero(number int) int {
	if number < 1 {
		return number
	}
	number = number - 1
	return countToZero(number)
}

func Benchmark2PowN(b *testing.B) {
	for _, test := range counttests {
		b.Run(fmt.Sprintf("input_%d", test.input), func(b *testing.B) {
			for count := 0; count < b.N; count++ {
				countToZero(test.input)
			}
		})
	}
}
