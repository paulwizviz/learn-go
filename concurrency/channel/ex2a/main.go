package main

import "fmt"

func main() {
	// Create a buffered channel
	c := make(chan string, 1)
	// Push a message to a buffered channel
	c <- "Hello"
	fmt.Println(<-c)
}
