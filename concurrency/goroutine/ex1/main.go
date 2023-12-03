package main

import (
	"fmt"
)

func main() {
	go func() {
		// This is a goroutine that is concurrent to the main routine
		fmt.Println("Hello") // This blocks the goroutine for a period of time. IO operations are typically blocked
	}()

	// This is the main routine and it executes immediately after the goroutine executes
	fmt.Println("World") // This operations blocks the main routine but it does not block long enough for the goroutine's print operations to execute
}
