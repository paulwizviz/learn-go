package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func(ch chan string) {
		fmt.Println("Start of goroutine")
		ch <- "Hello" // This channel is blocked
		fmt.Println("End of goroutine")
	}(c)

	fmt.Println("Put main routine to sleep")
	time.Sleep(10 * time.Second)
	fmt.Println("Incoming message: ", <-c)
}

// Before main routine awakes, you will see this in concole:
//   Put main routine to sleep
//   Start of goroutine
// When the main routine awakes, you will see this in console:
//   Put main routine to sleep
//   Start of goroutine
//   End of goroutine
//   Incoming message:  Hello
