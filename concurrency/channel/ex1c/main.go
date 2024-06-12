package main

import (
	"fmt"
	"log"
)

func main() {
	c := make(chan string)
	// Goroutine execute and handover to main routine
	go func(ch chan<- string) {
		// Send message to channel
		for range 2 {
			c <- "Hello World"
		}
		close(c) // Channel is closed no more message allowed and signal to receiver that no more is coming.
	}(c)
	// We have an endless loop but no idea of the number of messages sent
	// or when the channel is closed.
	for {
		// Receive message
		v, ok := <-c
		// If channel is closed ok returns false otherwise channel is opened and expect more messages
		if !ok {
			log.Fatal("Channel closed")
		}
		fmt.Println(v)
	}

}
