package bit

import "fmt"

func Example_binary() {
	binaryOne := 0b001
	binaryTwo := 0b010

	fmt.Printf("bits - One: %b Two: %b\n", binaryOne, binaryTwo)
	fmt.Printf("3 bits - One: %03b Two: %03b\n", binaryOne, binaryTwo)
	fmt.Printf("8 bits - One: %08b Two: %08b\n", binaryOne, binaryTwo)
	fmt.Printf("Decimal - One: %v Two: %v\n", binaryOne, binaryTwo)

	// Output:
	// bits - One: 1 Two: 10
	// 3 bits - One: 001 Two: 010
	// 8 bits - One: 00000001 Two: 00000010
	// Decimal - One: 1 Two: 2
}

func Example_octal() {
	octal0 := 00
	octal1 := 01
	octal2 := 02
	octal3 := 03
	octal4 := 04
	octal5 := 05
	octal6 := 06
	octal7 := 07
	octal10 := 010

	fmt.Printf("Octal0 %v\n", octal0)
	fmt.Printf("Octal1 %v\n", octal1)
	fmt.Printf("Octal2 %v\n", octal2)
	fmt.Printf("Octal3 %v\n", octal3)
	fmt.Printf("Octal4 %v\n", octal4)
	fmt.Printf("Octal5 %v\n", octal5)
	fmt.Printf("Octal6 %v\n", octal6)
	fmt.Printf("Octal7 %v\n", octal7)
	fmt.Printf("Octal10 %v\n", octal10)

	fmt.Printf("Octal0 Literal %0o\n", octal0)
	fmt.Printf("Octal1 Literal %0o\n", octal1)
	fmt.Printf("Octal2 Literal %0o\n", octal2)
	fmt.Printf("Octal3 Literal %0o\n", octal3)
	fmt.Printf("Octal4 Literal %0o\n", octal4)
	fmt.Printf("Octal5 Literal %0o\n", octal5)
	fmt.Printf("Octal6 Literal %0o\n", octal6)
	fmt.Printf("Octal7 Literal %0o\n", octal7)
	fmt.Printf("Octal10 Literal %0o", octal10)

	// Output:
	// Octal0 0
	// Octal1 1
	// Octal2 2
	// Octal3 3
	// Octal4 4
	// Octal5 5
	// Octal6 6
	// Octal7 7
	// Octal10 8
	// Octal0 Literal 0
	// Octal1 Literal 1
	// Octal2 Literal 2
	// Octal3 Literal 3
	// Octal4 Literal 4
	// Octal5 Literal 5
	// Octal6 Literal 6
	// Octal7 Literal 7
	// Octal10 Literal 10
}

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

	fmt.Printf("Hex00 in string %x\n", hex00)
	fmt.Printf("Hex01 in string %x\n", hex01)
	fmt.Printf("Hex02 in string %x\n", hex02)
	fmt.Printf("Hex03 in string %x\n", hex03)
	fmt.Printf("Hex04 in string %x\n", hex04)
	fmt.Printf("Hex05 in string %x\n", hex05)
	fmt.Printf("Hex06 in string %x\n", hex06)
	fmt.Printf("Hex07 in string %x\n", hex07)
	fmt.Printf("Hex08 in string %x\n", hex08)
	fmt.Printf("Hex09 in string %x\n", hex09)
	fmt.Printf("Hex0A in string %x\n", hex0A)
	fmt.Printf("Hex0B in string %x\n", hex0B)
	fmt.Printf("Hex0C in string %x\n", hex0C)
	fmt.Printf("Hex0D in string %x\n", hex0D)
	fmt.Printf("Hex0E in string %x\n", hex0E)
	fmt.Printf("Hex0F in string %x\n", hex0F)
	fmt.Printf("Hex0F in string %X\n", hex0F)

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
	// Hex00 in string 0
	// Hex01 in string 1
	// Hex02 in string 2
	// Hex03 in string 3
	// Hex04 in string 4
	// Hex05 in string 5
	// Hex06 in string 6
	// Hex07 in string 7
	// Hex08 in string 8
	// Hex09 in string 9
	// Hex0A in string a
	// Hex0B in string b
	// Hex0C in string c
	// Hex0D in string d
	// Hex0E in string e
	// Hex0F in string f
	// Hex0F in string F
}

func Example_bytes() {

	byte1 := byte(1)
	byte2 := byte(2)
	byte100 := byte(100)
	byte255 := byte(255)
	// byte := byte(256) is out of range

	fmt.Printf("byte1 %v\n", byte1)
	fmt.Printf("byte2 %v\n", byte2)
	fmt.Printf("byte100 %v\n", byte100)

	fmt.Printf("byte1 in binary %08b\n", byte1)
	fmt.Printf("byte2 in binary %08b\n", byte2)
	fmt.Printf("byte255 in binary %v\n", byte255)

	// Output:
	// byte1 1
	// byte2 2
	// byte100 100
	// byte1 in binary 00000001
	// byte2 in binary 00000010
	// byte255 in binary 255
}
