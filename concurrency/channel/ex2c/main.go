package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2) // Create a channel with two slices of strings
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

	c <- "Hello"
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

	c <- "World"
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

	c <- "Hola" // Send a message to channel when its buffer is full this will cause a deadlock as there is no receiving channel to drain the buffer.
	// fatal error: all goroutines are asleep - deadlock!
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

}
