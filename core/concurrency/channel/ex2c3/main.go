package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)
	go func(ch chan int) {
		defer close(c)
		for i := range 10 {
			ch <- i
		}
	}(c)
	for {
		time.Sleep(1 * time.Second)
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Println("Value: ", v)
	}
	fmt.Println("End")
}
