package types

import (
	"fmt"
	"math"
)

func Example_int() {

	// int is machine dependent so to get max and min value use the math package
	fmt.Printf("Smallest int %v\n", math.MinInt)
	fmt.Printf("Max int %v\n", math.MaxInt)

	var maxuint uint = math.MaxUint
	fmt.Printf("Max uint %v\n", maxuint)
	// fmt.Printf("Max uint %d\n", math.MaxUint) - This causes overflow

	// Output:
	// Smallest int -9223372036854775808
	// Max int 9223372036854775807
	// Max uint 18446744073709551615
}

func Example_uint8() {

	var minusi8 int8 = -128 // -129 will cause overflow
	var maxi8 int8 = 127    // 128 will cause overflow
	var u8 uint8 = 255      // -1 and 256 will cause overflow

	fmt.Printf("Smallest value i8 %v\n", minusi8)
	fmt.Printf("Max value for i8 %v\n", maxi8)
	fmt.Printf("Max value for u8 %v\n", u8)

	fmt.Printf("Math package int8. Min: %v Max: %v\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("Math package uint8. Max: %v", math.MaxUint8)

	// Output:
	// Smallest value i8 -128
	// Max value for i8 127
	// Max value for u8 255
	// Math package int8. Min: -128 Max: 127
	// Math package uint8. Max: 255

}

func Example_int16() {
	var minusi16 int16 = -32_768 // -32,769  will cause overflow
	var maxi16 int16 = 32_767    // 32,768 will cause overflow
	var u16 uint16 = 65_535      // 65,535 will cause overflow

	fmt.Printf("Smallest value i16 %v\n", minusi16)
	fmt.Printf("Maximum value i16 %v\n", maxi16)
	fmt.Printf("Maximum value u16 %v\n", u16)

	// Output:
	// Smallest value i16 -32768
	// Maximum value i16 32767
	// Maximum value u16 65535

}

func Example_int32() {
	var minusi32 int32 = -2_147_483_648 // -2,147,483,649  will cause overflow
	var maxi32 int32 = 2_147_483_647    // 2,147,483,648 will cause overflow
	var u32 uint32 = 4_294_967_295      // 4,294,967,296 will cause overflow

	fmt.Printf("Smallest value i32 %v\n", minusi32)
	fmt.Printf("Maximum value i32 %v\n", maxi32)
	fmt.Printf("Maximum value u32 %v\n", u32)

	// Output:
	// Smallest value i32 -2147483648
	// Maximum value i32 2147483647
	// Maximum value u32 4294967295

}

func Example_int64() {
	var minusi64 int64 = -922_337_203_685_4775_808 // -922_337_203_685_4775_809  will cause overflow
	var maxi64 int64 = 922_337_203_685_477_5807    // 922_337_203_685_477_5807 will cause overflow
	var u64 uint64 = 18_446_744_073_709_551_615    // 18_446_744_073_709_551_616 will cause overflow

	fmt.Printf("Smallest value i64 %v\n", minusi64)
	fmt.Printf("Maximum value i64 %v\n", maxi64)
	fmt.Printf("Maximum value u64 %v\n", u64)

	// Output:
	// Smallest value i64 -9223372036854775808
	// Maximum value i64 9223372036854775807
	// Maximum value u64 18446744073709551615
}

func Example_printInt() {
	i := 8

	fmt.Printf("%d\n", i)
	fmt.Printf("%8d\n", i)
	fmt.Printf("%08d", i)

	// Output:
	// 8
	//        8
	// 00000008
}
