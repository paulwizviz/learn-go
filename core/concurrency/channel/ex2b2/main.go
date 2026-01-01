package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 2)

	go func(ch chan int) {
		time.Sleep(10 * time.Second)
		fmt.Println("Push to channel")
		ch <- 1
	}(c)

	fmt.Println("Waiting")
	<-c
	time.Sleep(time.Second)
	fmt.Println("End")
}
