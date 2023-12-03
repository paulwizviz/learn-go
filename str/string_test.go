package str

import "fmt"

func Example_literal() {

	raw := `\n`
	interpreted := "\n"
	hello := "Hello"
	thumbsYellowUp := "👍"
	thumbsBrownUp := "👍🏾" //UTF-8 multiple bytes

	fmt.Println("Raw: ", []byte(raw))
	fmt.Println("Interpreted: ", []byte(interpreted))
	// len returns the number of bytes
	fmt.Println("Runes: ", []rune(hello), "Bytes: ", []byte(hello), "Byte size: ", len(hello))
	fmt.Println("Yellow thumbs up runes: ", []rune(thumbsYellowUp), "Length: ", len([]rune(thumbsYellowUp)))
	fmt.Println("Yellow thumbs up length: ", len(thumbsYellowUp), "Bytes: ", []byte(thumbsYellowUp))
	fmt.Println("Brown thumbs up runes: ", []rune(thumbsBrownUp), "Length: ", len([]rune(thumbsBrownUp)))
	fmt.Println("Brown thumbs up length: ", len(thumbsBrownUp), "Bytes: ", []byte(thumbsBrownUp))

	// Output:
	// Raw:  [92 110]
	// Interpreted:  [10]
	// Runes:  [72 101 108 108 111] Bytes:  [72 101 108 108 111] Byte size:  5
	// Yellow thumbs up runes:  [128077] Length:  1
	// Yellow thumbs up length:  4 Bytes:  [240 159 145 141]
	// Brown thumbs up runes:  [128077 127998] Length:  2
	// Brown thumbs up length:  8 Bytes:  [240 159 145 141 240 159 143 190]

}

func Example_utf8() {
	str := "日" // 3 byte character
	fmt.Printf("Length: %v Bytes: %v\n", len(str), []byte(str))
	fmt.Printf("Index 0 Character: %c\n", str[0])
	fmt.Printf("Index 1 Character: %c\n", str[1])
	fmt.Printf("Index 2 Character: %c\n", str[2])
	fmt.Printf("Character: %v", str[0:3])

	// Output:
	// Length: 3 Bytes: [230 151 165]
	// Index 0 Character: æ
	// Index 1 Character: 
	// Index 2 Character: ¥
	// Character: 日
}

func Example_ranging() {
	str := "a日bc" // This string includes multibyte character
	fmt.Printf("Character: %c Bytes: %v Length: %v\n", str[3], []byte(str), len(str))

	fmt.Println("---")

	for i, s := range str {
		fmt.Printf("Index: %d, Character: %c\n", i, s)
	}

	fmt.Println("---")

	for i, b := range []byte(str) {
		fmt.Printf("Index: %d, Character: %c\n", i, b)
	}

	fmt.Println("---")

	fmt.Printf("Character: %v", str[1:4])

	// Output:
	// Character: ¥ Bytes: [97 230 151 165 98 99] Length: 6
	// ---
	// Index: 0, Character: a
	// Index: 1, Character: 日
	// Index: 4, Character: b
	// Index: 5, Character: c
	// ---
	// Index: 0, Character: a
	// Index: 1, Character: æ
	// Index: 2, Character: 
	// Index: 3, Character: ¥
	// Index: 4, Character: b
	// Index: 5, Character: c
	// ---
	// Character: 日
}
