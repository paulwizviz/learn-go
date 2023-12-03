package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func(ch chan string) {
		time.Sleep(10 * time.Millisecond)
		ch <- "Sleep for 10 ms"
	}(c1)

	go func(ch chan string) {
		time.Sleep(1 * time.Millisecond)
		ch <- "Sleep for 1 ms"
	}(c2)

	c3 := time.After(20 * time.Millisecond) // This function returns a channel

	for {
		select {
		case msg := <-c1: // This channel blocks the main routine for 10 ms
			fmt.Println(msg) // This prints next
		case msg := <-c2: // This channel blocks the main routine for 1 ms
			fmt.Println(msg) // This prints first
		case <-c3: // This channel blocks the main routine for 20 ms
			fmt.Println("Timeout") // This prints last
			return
		}
	}
}
