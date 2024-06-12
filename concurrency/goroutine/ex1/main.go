package main

import (
	"fmt"
)

func main() {
	go func() {
		// This is operation executes in a goroutine
		// it runs and than hand off to the main routine
		fmt.Println("Hello") // io operations. It requires time to complete execution
	}()

	// This is the main rountine it ends before go routine ends
	// Hence no hello
}
