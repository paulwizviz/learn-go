package main

import "fmt"

func main() {
	c := make(chan string)
	go func(ch chan<- string) {
		defer close(ch) // Close channel at the end of the goroutine
		ch <- "Hello"
	}(c)

	v, opened := <-c // receive message before it is closed
	if opened {
		fmt.Printf("The channel is opened: %v\n", opened) // opened is true
		fmt.Printf("Message is: %v\n", v)
	}

	v, opened = <-c
	if !opened {
		fmt.Printf("The channel is opened: %v\n", opened) // opened is false
		fmt.Printf("Message is: %v\n", v)
	}
}
