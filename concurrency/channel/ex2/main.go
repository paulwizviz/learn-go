package main

import "fmt"

func main() {
	c := make(chan string)
	go func(ch chan<- string) {
		for i := 1; i < 6; i++ {
			ch <- "hello"
		}
	}(c)
	for {
		fmt.Println(<-c) // After five incoming messagevs you get a deadlock and causing the main routine to panic
	}
}
