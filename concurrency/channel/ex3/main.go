package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func(ch chan<- string) {
		fmt.Println("Hello world") // 1 - This block the goroutine
		ch <- "Hello"              // 3 - This is sent after channel is closed causing the routine to panic
	}(c)
	close(c) // 2 - Whilst the goroutine is block, this closes before the message hello is sent
	time.Sleep(1 * time.Millisecond)
}
