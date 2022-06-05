package slices_test

import (
	"fmt"
)

func Example_indexArray() {

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

	fmt.Println(data[DEC])

	calendar := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}
	fmt.Printf("len(calendar): %v cap(calendar): %v calendar[0]: %v calendar[1]: %v", len(calendar), cap(calendar), calendar[0], calendar[1])

	// Output:
	// December
	// len(calendar): 13 cap(calendar): 13 calendar[0]:  calendar[1]: Jan
}

func Example_array() {
	data := [3]int{1, 2}
	output := fmt.Sprintf("Type: %T Value: %v Length: %v Capacity: %v", data, data, len(data), cap(data))
	fmt.Println(output)

	data[1] = 1
	fmt.Println(data)

	/*
		data = append(data, 1) // Build error: first argument to append must be slice; have [3]int
	*/

	// Output:
	// Type: [3]int Value: [1 2 0] Length: 3 Capacity: 3
	// [1 1 0]
}

func Example_makeArray() {
	fixCapacityArray := make([]int, 10)
	fmt.Printf("Length: %v Capacity: %v Data: %v\n", len(fixCapacityArray), cap(fixCapacityArray), fixCapacityArray)

	fixCapacityArray = append(fixCapacityArray, 1)
	fmt.Printf("Data: %v\n", fixCapacityArray)

	fixLengthArray := make([]int, 1, 10)
	fmt.Printf("Length: %v Capacity: %v Data: %v\n", len(fixLengthArray), cap(fixLengthArray), fixLengthArray)

	fixLengthArray[0] = 1
	fmt.Printf("Length: %v Capacity: %v Data: %v\n", len(fixLengthArray), cap(fixLengthArray), fixLengthArray)

	/*
		fixLengthArray[1] = 1
		fmt.Printf("Length: %v Capacity: %v Data: %v\n", len(fixLengthArray), cap(fixLengthArray), fixLengthArray) // panic: runtime error: index out of range [1] with length 1
	*/

	fixLengthArray = append(fixLengthArray, 2)
	fmt.Printf("Length: %v Capacity: %v Data: %v\n", len(fixLengthArray), cap(fixLengthArray), fixLengthArray)

	// Output:
	// Length: 10 Capacity: 10 Data: [0 0 0 0 0 0 0 0 0 0]
	// Data: [0 0 0 0 0 0 0 0 0 0 1]
	// Length: 1 Capacity: 10 Data: [0]
	// Length: 1 Capacity: 10 Data: [1]
	// Length: 2 Capacity: 10 Data: [1 2]
}
