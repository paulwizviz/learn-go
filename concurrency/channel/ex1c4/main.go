package main

import (
	"fmt"
	"log"
)

func main() {
	c := make(chan string)
	s := make(chan struct{})
	go func(ch chan<- string, sig chan<- struct{}) {
		// Sends five messages then stops
		// sending to channel
		for i := 1; i < 6; i++ {
			ch <- "hello"
		}
		// We send an arbitrary signal to quit
		// This is used in lieu of close.
		sig <- struct{}{}
	}(c, s)
	for {
		// Using select to protect message
		select {
		case m := <-c:
			fmt.Println(m)
		case <-s:
			log.Fatal("End of messages")
		}
	}
}
