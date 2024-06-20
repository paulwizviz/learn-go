package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)
	go func(ch chan int) {
		fmt.Println("Fill buffer with 2 message")
		ch <- 1
		ch <- 2
		fmt.Println("Buffer is full")
		ch <- 3 // This is blocked until the buffer is drained
		fmt.Println("End of goroutine")
	}(c)

	fmt.Println("Put main routine to sleep")
	time.Sleep(10 * time.Second)
	fmt.Println("First message: ", <-c)
	fmt.Println("Second message: ", <-c)
	fmt.Println("Put main routine to sleep again")
	time.Sleep(10 * time.Second)
	fmt.Println("Third message: ", <-c) // Pull more message after sleep.

}
