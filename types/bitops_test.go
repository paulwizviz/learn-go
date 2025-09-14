package types

import "fmt"

func Example_bitwiseAND() {
	// & bitwise AND

	fmt.Printf("0b01 & 0b01 = %x\n", 0b01&0b01)
	fmt.Printf("0b01 & 0b00 = %x\n", 0b01&0b00)
	fmt.Printf("0b00 & 0b00 = %x\n", 0b00&0b00)

	// Output:
	// 0b01 & 0b01 = 1
	// 0b01 & 0b00 = 0
	// 0b00 & 0b00 = 0
}

func Example_bitwiseOR() {

	// | bitwise OR
	fmt.Printf("0b01 | 0b01 = %x\n", 0b01&0b01)
	fmt.Printf("0b01 | 0b00 = %x\n", 0b01|0b00)
	fmt.Printf("0b00 | 0b00 = %x\n", 0b00|0b00)

	// Output:
	// 0b01 | 0b01 = 1
	// 0b01 | 0b00 = 1
	// 0b00 | 0b00 = 0
}

func Example_bitwiseXOR() {

	// ^ bitwise XOR
	fmt.Printf("0b01 ^ 0b01 = %x\n", 0b01^0b01)
	fmt.Printf("0b01 ^ 0b00 = %x\n", 0b01^0b00)
	fmt.Printf("0b00 ^ 0b00 = %x\n", 0b00^0b00)

	// Output:
	// 0b01 ^ 0b01 = 0
	// 0b01 ^ 0b00 = 1
	// 0b00 ^ 0b00 = 0
}

func Example_andnot() {

	// &^   AND NOT

	fmt.Printf("0b01 &^ 0b01 = %x\n", 0b01&^0b01)
	fmt.Printf("0b01 &^ 0b00 = %x\n", 0b01&^0b00)
	fmt.Printf("0b00 &^ 0b00 = %x\n", 0b00&^0b00)

	// Output:
	// 0b01 &^ 0b01 = 0
	// 0b01 &^ 0b00 = 1
	// 0b00 &^ 0b00 = 0

}

func Example_bitshift() {

	// <<   left shift
	// >>   right shift

	fmt.Printf("0x0F>>1 = %08b\n", 0x0F>>1)
	fmt.Printf("0x0F<<1 = %08b\n", 0x0F<<1)
	fmt.Printf("0x0F>>1 = %v\n", 0x0F>>1)
	fmt.Printf("0x0F<<1 = %v\n", 0x0F<<1)

	// Output:
	// 0x0F>>1 = 00000111
	// 0x0F<<1 = 00011110
	// 0x0F>>1 = 7
	// 0x0F<<1 = 30
}

// Using bitwise operations to determine odds and even
func Example_oddeven() {

	// odd & 1 returns 1
	odd := 115
	fmt.Println(fmt.Sprintf("Odd number: %d", odd), fmt.Sprintf("odd & 1 = %d", odd&1))

	// even & 1 returns 0
	even := 114
	fmt.Println(fmt.Sprintf("Even number: %d", even), fmt.Sprintf("even & 1 = %d", even&1))

	// Output:
	// Odd number: 115 odd & 1 = 1
	// Even number: 114 even & 1 = 0
}
