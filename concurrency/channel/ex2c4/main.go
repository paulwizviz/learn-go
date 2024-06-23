package main

import "fmt"

func asyncProcessor() chan int {
	c := make(chan int, 1)
	go func(ch chan int) {
		for i := range 10 {
			ch <- i
		}
		close(ch)
	}(c)
	return c
}

func main() {
	data := asyncProcessor()
	for d := range data {
		fmt.Println(d)
	}
}
