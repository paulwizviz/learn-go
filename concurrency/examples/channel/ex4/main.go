package main

import "fmt"

func main() {
	c := make(chan string)
	go func(ch chan<- string) {
		defer close(ch)
		ch <- "Hello"
	}(c)

	v, opened := <-c
	if opened {
		fmt.Printf("The channel is opened: %v\n", opened)
		fmt.Printf("Message is: %v\n", v)
	}

	v, opened = <-c
	if !opened {
		fmt.Printf("The channel is opened: %v\n", opened)
		fmt.Printf("Message is: %v\n", v)
	}
}
