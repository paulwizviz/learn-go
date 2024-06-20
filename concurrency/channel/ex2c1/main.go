package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	go func(ch chan int) {
		for i := range 10 {
			ch <- i
		}
	}(c)
	for {
		fmt.Println("Message: ", <-c)
	}
}
