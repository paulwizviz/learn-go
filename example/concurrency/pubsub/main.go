package main

import (
	"fmt"
)

type broker struct {
	subscribers []chan int
}

func publishTo(b *broker, d int) {
	for _, ch := range b.subscribers {
		ch <- d
	}
}

func subscribeTo(b *broker, ch chan int) {
	b.subscribers = append(b.subscribers, ch)
}

func main() {

	dataset := []int{10, 20, 40, 100}

	b := &broker{}

	c1 := make(chan int, len(dataset))
	c2 := make(chan int, len(dataset))

	subscribeTo(b, c1)
	subscribeTo(b, c2)

	go func() {
		for _, d := range dataset {
			publishTo(b, d)
		}
		for i := 0; i < len(b.subscribers); i++ {
			close(b.subscribers[i])
		}
	}()

	for c := range c1 {
		fmt.Println("Result for subscriber 1", c)
	}

	for c := range c2 {
		fmt.Println("Result for subscriber 2", c)
	}

}
