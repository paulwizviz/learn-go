package slices

import (
	"fmt"
)

func Example_slice() {
	emptySlice := []int{}
	fmt.Printf("Capacity: %v Length: %v Data: %v\n", cap(emptySlice), len(emptySlice), emptySlice)

	emptySlice = append(emptySlice, 1)
	fmt.Printf("Capacity: %v Length: %v Data: %v\n", cap(emptySlice), len(emptySlice), emptySlice)

	nilSlice := []int(nil)
	fmt.Printf("Capacity: %v Length: %v Data: %v\n", cap(nilSlice), len(nilSlice), nilSlice)

	var arr [4]string
	strSlice := []string{"abc", "efg"}
	copy(arr[:], strSlice[:])
	fmt.Printf("Type: %T Data: %v Type: %T Data: %v\n", arr, arr, strSlice, strSlice)

	// Output:
	// Capacity: 0 Length: 0 Data: []
	// Capacity: 1 Length: 1 Data: [1]
	// Capacity: 0 Length: 0 Data: []
	// Type: [4]string Data: [abc efg  ] Type: []string Data: [abc efg]
}

func Example_indexSlice() {

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

func ExampleReverseInt() {
	input := []int{1, 2, 3, 4, 5}
	output := ReverseInt(input)
	fmt.Println(output)
	// Output:
	// [5 4 3 2 1]
}
