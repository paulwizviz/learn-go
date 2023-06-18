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

	c3 := time.After(20 * time.Millisecond)

	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		case <-c3:
			fmt.Println("Timeout")
			return
		}
	}
}
