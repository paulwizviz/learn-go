package types

import "fmt"

func Example_rune() {
	a := 'a'
	specialChar := '\377'
	chineseDay := '日'
	lowerADecimal := 97
	lowerAOctal := '\141'
	lowerAHex := '\x61'
	lowerAuc1 := '\u0061'     // unicode
	lowerAuc2 := '\U00000061' // utf-8

	fmt.Printf("Type: %T Value: %d Graphine: %c\n", a, a, a)
	fmt.Printf("Type: %T Value: %d Graphine: %c\n", specialChar, specialChar, specialChar)
	fmt.Printf("Type: %T Value: %d Graphine: %c\n", chineseDay, chineseDay, chineseDay)
	fmt.Printf("Type: %T Value: %d Graphine: %c\n", lowerADecimal, lowerADecimal, lowerADecimal)
	fmt.Printf("Type: %T Value: %d Graphine: %c\n", lowerAOctal, lowerAOctal, lowerAOctal)
	fmt.Printf("Type: %T Value: %d Graphine: %c\n", lowerAHex, lowerAHex, lowerAHex)
	fmt.Printf("Type: %T Value: %d Graphine: %c\n", lowerAuc1, lowerAuc1, lowerAuc1)
	fmt.Printf("Type: %T Value: %d Graphine: %c\n", lowerAuc2, lowerAuc2, lowerAuc2)

	// Output:
	// Type: int32 Value: 97 Graphine: a
	// Type: int32 Value: 255 Graphine: ÿ
	// Type: int32 Value: 26085 Graphine: 日
	// Type: int Value: 97 Graphine: a
	// Type: int32 Value: 97 Graphine: a
	// Type: int32 Value: 97 Graphine: a
	// Type: int32 Value: 97 Graphine: a
	// Type: int32 Value: 97 Graphine: a
}
