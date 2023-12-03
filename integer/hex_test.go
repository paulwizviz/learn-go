package integer

import "fmt"

func Example_hex() {

	hex00 := 0x00
	hex01 := 0x01
	hex02 := 0x02
	hex03 := 0x03
	hex04 := 0x04
	hex05 := 0x05
	hex06 := 0x06
	hex07 := 0x07
	hex08 := 0x08
	hex09 := 0x09
	hex0A := 0x0A
	hex0B := 0x0B
	hex0C := 0x0C
	hex0D := 0x0D
	hex0E := 0x0E
	hex0F := 0x0F

	fmt.Printf("Hex00 in decimal %v\n", hex00)
	fmt.Printf("Hex01 in decimal %v\n", hex01)
	fmt.Printf("Hex02 in decimal %v\n", hex02)
	fmt.Printf("Hex03 in decimal %v\n", hex03)
	fmt.Printf("Hex04 in decimal %v\n", hex04)
	fmt.Printf("Hex05 in decimal %v\n", hex05)
	fmt.Printf("Hex06 in decimal %v\n", hex06)
	fmt.Printf("Hex07 in decimal %v\n", hex07)
	fmt.Printf("Hex08 in decimal %v\n", hex08)
	fmt.Printf("Hex09 in decimal %v\n", hex09)
	fmt.Printf("Hex0A in decimal %v\n", hex0A)
	fmt.Printf("Hex0B in decimal %v\n", hex0B)
	fmt.Printf("Hex0C in decimal %v\n", hex0C)
	fmt.Printf("Hex0D in decimal %v\n", hex0D)
	fmt.Printf("Hex0E in decimal %v\n", hex0E)
	fmt.Printf("Hex0F in decimal %v\n", hex0F)

	// Output:
	// Hex00 in decimal 0
	// Hex01 in decimal 1
	// Hex02 in decimal 2
	// Hex03 in decimal 3
	// Hex04 in decimal 4
	// Hex05 in decimal 5
	// Hex06 in decimal 6
	// Hex07 in decimal 7
	// Hex08 in decimal 8
	// Hex09 in decimal 9
	// Hex0A in decimal 10
	// Hex0B in decimal 11
	// Hex0C in decimal 12
	// Hex0D in decimal 13
	// Hex0E in decimal 14
	// Hex0F in decimal 15
}

func Example_hexint() {
	var hex1 uint8 = 0xFF
	var hex2 int8 = 0x7E
	var hex3 int8 = -0x80

	fmt.Printf("Max uint8 Hex: %X Decimal: %d\n", hex1, hex1)
	fmt.Printf("Max int8 Hex: %X Decimal: %d\n", hex2, hex2)
	fmt.Printf("Min int8 Hex: %X Decimal: %d\n", hex3, hex3)

	// Output:
	// Max uint8 Hex: FF Decimal: 255
	// Max int8 Hex: 7E Decimal: 126
	// Min int8 Hex: -80 Decimal: -128
}

func Example_printhex() {

	i := 11
	h := 0xABCDEF

	fmt.Printf("Lower: %x Upper: %X\n", i, i)
	fmt.Printf("Lower: %x Upper: %X\n", h, h)
	fmt.Printf("%8X\n", h)
	fmt.Printf("%08X", h)

	// Output:
	// Lower: b Upper: B
	// Lower: abcdef Upper: ABCDEF
	//   ABCDEF
	// 00ABCDEF
}
