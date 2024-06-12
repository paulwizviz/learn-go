package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func(c chan string) {
		defer close(c)
		for {
			time.Sleep(10 * time.Second)
			c <- "Hello"
			break
		}
	}(c)
	fmt.Println("Wait for message")
	fmt.Println(<-c)
	fmt.Println("End of message")
}
