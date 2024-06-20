package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func(ch chan string) {
		fmt.Println("Start of goroutine")
		time.Sleep(10 * time.Second)
		ch <- "Hello"
		fmt.Println("End of goroutine")
	}(c)

	fmt.Println("Waiting to receive message from channel")
	fmt.Println("Incoming message: ", <-c) // This is blocked until sender channel has received message
}
