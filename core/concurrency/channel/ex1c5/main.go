package main

import "fmt"

func main() {
	c := make(chan string)
	go func(ch chan<- string) {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- "Hello"
		}
	}(c)

	for v := range c { // This ensure the receving channel range until all messages from sender are drained
		fmt.Println(v)
	}
}
