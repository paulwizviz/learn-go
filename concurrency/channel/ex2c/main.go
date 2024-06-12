package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c) // This is trying to read more than what is available in the buffer
}
