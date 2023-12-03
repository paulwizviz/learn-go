package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)
	go func(ch chan int) {
		defer close(c)
		for i := 0; i < 10; i++ {
			select {
			case ch <- i:
				fmt.Printf("Pushing %v to channel\n", i)
			default:
				// This is called when the channel buffer is full.
				fmt.Printf("Don't push %v to channel\n", i)
			}
		}
	}(c)
	time.Sleep(10 * time.Millisecond)
}
