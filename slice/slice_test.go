package slice

import (
	"fmt"
	"slices"
)

func Example_initSlice() {
	s := []int{}
	fmt.Printf("Capacity: %v Length: %v Data: %v\n", cap(s), len(s), s)

	s = []int(nil)
	fmt.Printf("Capacity: %v Length: %v Data: %v\n", cap(s), len(s), s)

	s = make([]int, 2)
	fmt.Printf("Length: %v Capacity: %v Data: %v\n", len(s), cap(s), s)

	s = make([]int, 0, 2)
	fmt.Printf("Length: %v Capacity: %v Data: %v\n", len(s), cap(s), s)

	// Output:
	// Capacity: 0 Length: 0 Data: []
	// Capacity: 0 Length: 0 Data: []
	// Length: 2 Capacity: 2 Data: [0 0]
	// Length: 0 Capacity: 2 Data: []
}

func Example_makeSlice() {

	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	s := make([]int, 2)
	fmt.Printf("Length: %v Capacity: %v Data: %v\n", len(s), cap(s), s)
	s[0] = 1
	s[1] = 1
	s[2] = 1

	// Output:
	// Length: 2 Capacity: 2 Data: [0 0]
	// runtime error: index out of range [2] with length 2

}

func Example_indexedSlice() {

	const (
		// JAN is eqivalemt to 0
		JAN = iota
		FEB
		MAR
		APR
		MAY
		JUN
		JUL
		AUG
		SEP
		OCT
		NOV
		DEC
	)

	slice := []string{
		JAN: "January",
		FEB: "Feburary",
		MAR: "March",
		APR: "April",
		MAY: "May",
		JUN: "June",
		JUL: "July",
		AUG: "August",
		SEP: "September",
		OCT: "October",
		NOV: "November",
		DEC: "December",
	}

	fmt.Println(slice[0:1])
	fmt.Println(slice[JAN:FEB])
	fmt.Println(slice[:AUG])
	fmt.Println(slice[MAR:])
	fmt.Println(slice[:])

	// Output:
	// [January]
	// [January]
	// [January Feburary March April May June July]
	// [March April May June July August September October November December]
	// [January Feburary March April May June July August September October November December]
}

func Example_append() {

	source := []int{1, 2, 3}
	fmt.Printf("Source: %v\n", source)
	target := []int{1, 2}
	target = append(target, source...)
	fmt.Printf("Target: %v\n", target)

	// Output:
	// Source: [1 2 3]
	// Target: [1 2 1 2 3]

}

func Example_subSlice() {
	source := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("len(source): %v, cap(source): %v\n", len(source), cap(source))
	fmt.Printf("source[:3] %v source[4:] %v\n", source[:3], source[4:])

	s := source[0:1:len(source)]
	fmt.Printf("len(s): %v cap(s) %v\n", len(s), cap(s))

	// Output:
	// len(source): 7, cap(source): 7
	// source[:3] [1 2 3] source[4:] [5 6 7]
	// len(s): 1 cap(s) 7
}

func Example_removeElementSlice() {
	source := []int{1, 2, 3, 4, 5, 6, 7}
	result := append(source[:3], source[4:]...)
	fmt.Printf("Using append: %v\n", result)
	result = slices.Delete(source, 3, 4)
	fmt.Printf("Using slice package: %v, %v\n", result, result[:len(result)-1])

	// Output:
	// Using append: [1 2 3 5 6 7]
}

func Example_multiDSlice() {

	numInnerSlice := 2
	capacityOfInnerSlice := 2

	twoDSlice := make([][]int, numInnerSlice) // Preallocate external slice
	for i := range 2 {
		twoDSlice[i] = make([]int, 0, capacityOfInnerSlice) // Preallocate inner slice
	}

	fmt.Printf("len(twoDSlice): %v, cap(twoDSlice): %v\n", len(twoDSlice), cap(twoDSlice))
	fmt.Printf("len(twoDSlice[0]): %v, cap(twoDSlice[0]): %v\n", len(twoDSlice[0]), cap(twoDSlice[0]))

	// Output:
	// len(twoDSlice): 2, cap(twoDSlice): 2
	// len(twoDSlice[0]): 0, cap(twoDSlice[0]): 2

}
