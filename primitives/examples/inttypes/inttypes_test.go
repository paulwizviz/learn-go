package inttypes_test

import "fmt"

func Example() {
	var signedArcInt int = -32
	fmt.Printf("Architecture dependent signed integer: %v\n", signedArcInt)

	var unSignedArcInt uint = 32
	fmt.Printf("Architecture dependent unsigned integer: %v\n", unSignedArcInt)

	var signedInt8 int8 = -128
	fmt.Printf("Signed int8: %v\n", signedInt8)

	var unSignedInt8 uint8 = 255
	fmt.Printf("Unsigned int8: %v\n", unSignedInt8)

	var SignedInt16 int16 = -32_768
	fmt.Printf("Signed int16: %v\n", SignedInt16)

	var UnsignedInt16 uint16 = 65_535
	fmt.Printf("Unsigned int16: %v\n", UnsignedInt16)

	var SignedInt32 int32 = -2_147_483_648
	fmt.Printf("Signed int32: %v\n", SignedInt32)

	var UnsignedInt32 uint32 = 4_294_967_295
	fmt.Printf("Unsigned int32: %v\n", UnsignedInt32)

	var SignedInt64 int64 = -9.223_372e+18
	fmt.Printf("Signed int64: %v\n", SignedInt64)

	var UnsignedInt64 uint64 = 1.8_446_744e+19
	fmt.Printf("Unsigned int64: %v\n", UnsignedInt64)

	var b uint8 = byte('a')
	fmt.Printf("byte is an alias for uint8: %v\n", b)

	var r int32 = '😊'
	fmt.Printf("rune is an alias for int32: %v\n", r)

	var rUTF8 int32 = '\uE414'
	fmt.Printf("rune is an alias for int32: %v\n", rUTF8)

	// Output:
	// Architecture dependent signed integer: -32
	// Architecture dependent unsigned integer: 32
	// Signed int8: -128
	// Unsigned int8: 255
	// Signed int16: -32768
	// Unsigned int16: 65535
	// Signed int32: -2147483648
	// Unsigned int32: 4294967295
	// Signed int64: -9223372000000000000
	// Unsigned int64: 18446744000000000000
	// byte is an alias for uint8: 97
	// rune is an alias for int32: 128522
	// rune is an alias for int32: 58388
}
