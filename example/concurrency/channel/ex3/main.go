package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func(ch chan<- string) {
		fmt.Println("Hello world")
		ch <- "Hello"
	}(c)
	close(c)
	time.Sleep(1 * time.Millisecond)
}
