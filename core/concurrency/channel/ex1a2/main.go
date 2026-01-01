package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go func(ch chan string) {
		ch <- "Hello"
	}(c)
	fmt.Println("Incoming message: ", <-c)
}
