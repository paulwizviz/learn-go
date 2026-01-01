package types

import "fmt"

func Example_bytes() {

	byte1 := 1
	byte2 := byte(2)
	byte100 := byte(100)
	byte255 := byte(255)
	// byte := byte(256) is out of range

	fmt.Printf("byte1 %v\n", byte1)
	fmt.Printf("byte2 %v\n", byte2)
	fmt.Printf("byte100 %v\n", byte100)
	fmt.Printf("byte255 %v\n", byte255)

	// Output:
	// byte1 1
	// byte2 2
	// byte100 100
	// byte255 255
}

func Example_bytesBinary() {

	byte1 := 1
	byte2 := byte(2)
	byte100 := byte(100)
	byte255 := byte(255)

	fmt.Printf("byte1 in binary %08b\n", byte1)
	fmt.Printf("byte2 in binary %08b\n", byte2)
	fmt.Printf("byte100 in binary %08b\n", byte100)
	fmt.Printf("byte255 in binary %v\n", byte255)

	// Output:
	// byte1 in binary 00000001
	// byte2 in binary 00000010
	// byte100 in binary 01100100
	// byte255 in binary 255
}

func Example_byteint() {

	var b1 uint8 = byte(255)
	var b2 uint8 = 0xFF
	var b3 uint8 = 0b11111111
	var b4 int = 256 // out of range

	fmt.Printf("Byte - Hex: %X Binary: %08b Deciaml: %d\n", b1, b1, b1)
	fmt.Printf("Byte - Hex: %X Binary: %08b Deciaml: %d\n", b2, b2, b2)
	fmt.Printf("Byte - Hex: %X Binary: %08b Deciaml: %d\n", b3, b3, b3)
	fmt.Printf("Byte - Hex: %X Binary: %08b Deciaml: %d", b4, b4, b4)

	// Output:
	// Byte - Hex: FF Binary: 11111111 Deciaml: 255
	// Byte - Hex: FF Binary: 11111111 Deciaml: 255
	// Byte - Hex: FF Binary: 11111111 Deciaml: 255
	// Byte - Hex: 100 Binary: 100000000 Deciaml: 256

}
