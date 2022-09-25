package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan string)

	go func(ch chan<- string) {
		time.Sleep(500 * time.Millisecond)
		ch <- "Hello world"
	}(c)

	timeBefore := time.Now()
	fmt.Println(timeBefore.Unix())
	msg := <-c // process is blocked until it receive message from channel after 500ms

	timeAfter := time.Now()
	fmt.Println(timeAfter.Unix())

	interval := time.Duration(500 * time.Millisecond)
	if interval < timeAfter.Sub(timeBefore) {
		fmt.Printf("Message: %s", msg) // This will print after 500ms
	}
}