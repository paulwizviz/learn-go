package main

import (
	_ "embed"
	"fmt"
)

var (

	//go:embed data/words.txt
	fruits []byte

	//go:embed data/lang.txt
	lang []byte
)

func Example() {
	fmt.Println(string(fruits))
	fmt.Println("----")
	fmt.Println(string(lang))

	// output:
	// This
	// is
	// a
	// test
	// ----
	// English
	// Mandarin
	// Malay
}
