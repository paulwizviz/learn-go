package str

import "fmt"

func Example_rune() {
	a := 'a'
	specialChar := '\377'
	chineseDay := '日'
	uniA := '\u0041' // unicode representation of A
	utf8 := '\x41'   // utf-8 hex representation of A

	fmt.Printf("Type: %T Value: %d\n", a, a)
	fmt.Printf("Type: %T Value: %d\n", specialChar, specialChar)
	fmt.Printf("Type: %T Value: %d\n", chineseDay, chineseDay)
	fmt.Printf("Type: %T Value: %d\n", uniA, uniA)
	fmt.Printf("Type: %T Value: %d\n", utf8, utf8)

	// Output:
	// Type: int32 Value: 97
	// Type: int32 Value: 255
	// Type: int32 Value: 26085
	// Type: int32 Value: 65
	// Type: int32 Value: 65
}
