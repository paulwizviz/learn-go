package integer

import (
	"fmt"
	"math"
)

func Example_binary() {
	var b uint = 0b0000_0010
	var b1 int = -0b0000_0010
	var b2 uint8 = 0b1111_1111
	var b3 int8 = -0b10000000
	var b4 int8 = 0b01111111

	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", b1)
	fmt.Printf("%d\n", b2)
	fmt.Printf("%d\n", b3)
	fmt.Printf("%d\n", b4)

	// Output:
	// 2
	// -2
	// 255
	// -128
	// 127
}

func Example_printBinary() {

	b := 0b001 // This mapped to int

	fmt.Printf("bits %b\n", b)
	fmt.Printf("4 bits %04b\n", b)
	fmt.Printf("8 bits %08b\n", b)
	fmt.Printf("Max Int16 %b\n", math.MaxInt16)
	fmt.Printf("Min Int16 %b", math.MinInt16)

	// Output:
	// bits 1
	// 4 bits 0001
	// 8 bits 00000001
	// Max Int16 111111111111111
	// Min Int16 -1000000000000000
}
