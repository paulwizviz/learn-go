package integer

import (
	"fmt"
	"math"
)

func Example_binaryprintformat() {

	b1 := 0b001 // This mapped to int
	b2 := 0b010

	fmt.Printf("bits - One: %b Two: %b\n", b1, b2)
	fmt.Printf("3 bits - One: %03b Two: %03b\n", b1, b2)
	fmt.Printf("8 bits - One: %08b Two: %08b\n", b1, b2)
	fmt.Printf("Decimal - One: %v Two: %v\n", b1, b2)

	// Output:
	// bits - One: 1 Two: 10
	// 3 bits - One: 001 Two: 010
	// 8 bits - One: 00000001 Two: 00000010
	// Decimal - One: 1 Two: 2
}

func Example_binary() {
	var b uint = 0b0000_0010
	var b1 int = -0b0000_0010
	var b2 uint8 = 0b1111_1111
	var b3 int8 = -0b10000000
	var b4 int8 = 0b01111111
	var b16 int16 = math.MaxInt16

	fmt.Printf("Size not specified: %b 8bits %08b Decimal %d\n", b, b, b)
	fmt.Printf("Size not specified: %b 8bits %08b Decimal %d\n", b1, b1, b1)
	fmt.Printf("%b\n", b2)
	fmt.Printf("%b\n", b3)
	fmt.Printf("%b\n", b4)
	fmt.Printf("%08b\n", b16)

	// Output:
	// Size not specified: 10 8bits 00000010 Decimal 2
	// Size not specified: -10 8bits -0000010 Decimal -2
	// 11111111
	// -10000000
	// 1111111
	// 111111111111111
}
