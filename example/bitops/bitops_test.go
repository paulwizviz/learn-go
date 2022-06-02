package bitops

import "fmt"

func Example_bitwiseAND() {
	// & bitwise AND

	fmt.Printf("0x01 & 0x01 = %x\n", 0x01&0x01)
	fmt.Printf("0x01 & 0x00 = %x\n", 0x01&0x00)
	fmt.Printf("0x00 & 0x00 = %x\n", 0x00&0x00)

	// Output:
	// 0x01 & 0x01 = 1
	// 0x01 & 0x00 = 0
	// 0x00 & 0x00 = 0
}

func Example_bitwiseOR() {
	// | bitwise OR

	fmt.Printf("0x01 | 0x01 = %x\n", 0x01|0x01)
	fmt.Printf("0x01 | 0x00 = %x\n", 0x01|0x00)
	fmt.Printf("0x00 | 0x00 = %x\n", 0x00|0x00)

	// Output:
	// 0x01 | 0x01 = 1
	// 0x01 | 0x00 = 1
	// 0x00 | 0x00 = 0
}

func Example_bitwiseXOR() {

	// ^ bitwise XOR
	fmt.Printf("0x01 ^ 0x01 = %x\n", 0x01^0x01)
	fmt.Printf("0x01 ^ 0x00 = %x\n", 0x01^0x00)
	fmt.Printf("0x00 ^ 0x00 = %x\n", 0x00^0x00)

	// Output:
	// 0x01 ^ 0x01 = 0
	// 0x01 ^ 0x00 = 1
	// 0x00 ^ 0x00 = 0
}

func Example_andnot() {

	// &^   AND NOT

	fmt.Printf("0x01 &^ 0x01 = %x\n", 0x01&^0x01)
	fmt.Printf("0x01 &^ 0x00 = %x\n", 0x01&^0x00)
	fmt.Printf("0x00 &^ 0x00 = %x\n", 0x00&^0x00)

	// Output:
	// 0x01 &^ 0x01 = 0
	// 0x01 &^ 0x00 = 1
	// 0x00 &^ 0x00 = 0

}

func Example_bitshift() {

	// <<   left shift
	// >>   right shift

	fmt.Printf("0x0F>>1 = %v\n", 0x0F>>1) // 1111 >> 1 = 0111 (7)
	fmt.Printf("0x0F<<1 = %v\n", 0x0F<<1) // 1111 << 1 = 1111 (30)

	// Output:
	// 0x0F>>1 = 7
	// 0x0F<<1 = 30
}

func Example_oddeven() {

	odd := 115
	fmt.Println(fmt.Sprintf("Odd number: %d", odd), fmt.Sprintf("odd & 1 = %d", odd&1))

	even := 114
	fmt.Println(fmt.Sprintf("Even number: %d", even), fmt.Sprintf("even & 1 = %d", even&1))

	// Output:
	// Odd number: 115 odd & 1 = 1
	// Even number: 114 even & 1 = 0
}
