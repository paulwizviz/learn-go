package main

import (
	"fmt"
	"log"
)

func main() {
	c := make(chan string)
	q := make(chan struct{})
	go func(ch chan<- string) {
		// Sends five messages then stops
		// sending to channel
		for i := 1; i < 6; i++ {
			ch <- "hello"
		}
		// We send an arbitrary signal to quit
		// This is used in lieu of close.
		q <- struct{}{}
	}(c)
	for {
		// Using select to protect message
		select {
		case m := <-c:
			fmt.Println(m)
		case <-q:
			log.Fatal("End of messages")
		}
	}
}
