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

	c <- "Hola"
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

}
