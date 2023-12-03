package integer

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

func Example_bytesint() {

	var b1 uint8 = byte(255)

	fmt.Printf("Byte - Hex: %X Binary: %08b Deciaml: %d", b1, b1, b1)

	// Output:
	// Byte - Hex: FF Binary: 11111111 Deciaml: 255

}
