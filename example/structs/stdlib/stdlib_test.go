package stdlib

import (
	"bytes"
	"fmt"
	"log"
)

type hello interface {
	Prefix() string
}

func displayPrefix(h hello) string {
	return h.Prefix()
}

// This example demostrate the use of a user defined
// interface that matches a standard lib method
// logger prefix
func Example_logger() {
	var b bytes.Buffer
	logger := log.New(&b, "test", 1)
	result := displayPrefix(logger) // Invoke the method Prefix()
	fmt.Println(result)

	// Output:
	// test
}
