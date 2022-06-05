package ioexample

import (
	"fmt"
	"strings"
)

func Example_stringReader() {

	input := "Hello World"
	r := strings.NewReader(input)

	b := make([]byte, 2)
	n, _ := r.Read(b)
	fmt.Printf("Number of bytes: %d Value: %s\n", n, string(b))

	n, _ = r.Read(b)
	fmt.Printf("Number of bytes: %d Value: %s", n, string(b))

	// Output:
	// Number of bytes: 2 Value: He
	// Number of bytes: 2 Value: ll
}

func Example_customReader() {

	input := "abcdefg"
	r := NewCustomReader([]byte(input))
	b := make([]byte, 2)
	r.Read(b)
	fmt.Println(string(b))
	r.Read(b)
	fmt.Println(string(b))

	// Output:
	// ab
	// cd

}
