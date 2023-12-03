package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello, 世界")
	}()

	time.Sleep(100 * time.Microsecond) // This cause the main routine to sleep for 100 microsec.
	// The main routine sleeps long enought for goroutine to complete execution so you will see "Hello, 世界"
}
