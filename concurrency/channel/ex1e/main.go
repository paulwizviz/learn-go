package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	// Goroutine starts executing and hand off to main routine
	go func(ch chan<- string) {
		fmt.Println("Hello world") // 1 - This block the goroutine before the next
		ch <- "Hello"              // 3 - This is sent after channel is closed causing the routine to panic
	}(c)
	close(c)                         // 2 - This closes the channel before the message hello is sent
	time.Sleep(1 * time.Millisecond) // The main routine sleeps to allow the Goroutine to execute
}
