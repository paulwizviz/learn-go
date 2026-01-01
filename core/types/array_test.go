package types

import (
	"fmt"
)

func Example_initArray() {
	data := [3]int{1, 2}
	output := fmt.Sprintf("Type: %T Value: %v Length: %v Capacity: %v", data, data, len(data), cap(data))
	fmt.Println(output)

	// Output:
	// Type: [3]int Value: [1 2 0] Length: 3 Capacity: 3
}

func Example_implicitLength() {
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

	data := [...]string{
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

	fmt.Printf("Len(data): %v, Cap(data): %v\n", len(data), cap(data))
	fmt.Printf("data[0]: %v, data[JAN]: %v, data[DEC]: %v\n", data[0], data[JAN], data[DEC])

	// Output:
	// Len(data): 12, Cap(data): 12
	// data[0]: January, data[JAN]: January, data[DEC]: December
}

func Example_indexArray() {

	calendar := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}
	fmt.Printf("len(calendar): %v cap(calendar): %v calendar[0]: %v calendar[1]: %v\n", len(calendar), cap(calendar), calendar[0], calendar[1])
	fmt.Printf("calendar[0] %v, calendar[1] %v\n", calendar[0], calendar[1])

	// Output:
	// len(calendar): 13 cap(calendar): 13 calendar[0]:  calendar[1]: Jan
	// calendar[0] , calendar[1] Jan
}

func Example_twoDimensional() {
	a := [2][2]int{{1, 2}, {1, 2}}
	fmt.Println(a)

	a = [2][2]int{
		0: {1, 2},
		1: {1, 2},
	}
	fmt.Printf("len (a): %v, cap (a): %v, a: %v\n", len(a), cap(a), a)

	fmt.Printf("a[1]: %v, a[1][1]: %v\n", a[1], a[1][1])

	a = [2][2]int{}
	for i := range 2 {
		for j := range 2 {
			a[i] = [2]int{i, j}
		}
	}
	fmt.Println(a)

	// Output:
	// [[1 2] [1 2]]
	// len (a): 2, cap (a): 2, a: [[1 2] [1 2]]
	// a[1]: [1 2], a[1][1]: 2
	// [[0 1] [1 1]]
}

func Example_threeD() {
	a := [2][2][2]int{{{1, 2}, {2, 1}}, {{1, 2}, {2, 1}}}
	fmt.Println(a)

	a = [2][2][2]int{
		0: {
			0: {1, 2},
			1: {2, 1},
		},
		1: {
			0: {1, 2},
			1: {2, 1},
		},
	}
	fmt.Println(a)

	// Output:
	// [[[1 2] [2 1]] [[1 2] [2 1]]]
	// [[[1 2] [2 1]] [[1 2] [2 1]]]
}

func Example_copyArray() {
	source := [3]int{1, 2, 3}
	fmt.Println(source)

	target := [3]int{}
	target = source
	target[1] = 9
	fmt.Println(target)

	// Output:
	// [1 2 3]
	// [1 9 3]
}
