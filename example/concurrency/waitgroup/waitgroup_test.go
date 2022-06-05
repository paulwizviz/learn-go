package waitgroup_test

import (
	"fmt"
	"sync"
	"time"
)

func Example_waitForTwoSameTime() {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1) // tells wg to consider only one concurrent process -- i.e. shortest
	go func(ch chan<- string) {
		time.Sleep(time.Millisecond * 500)
		ch <- "Hello"
		wg.Done()
	}(c)

	go func(ch chan<- string) {
		time.Sleep(time.Millisecond * 500)
		ch <- "World"
		wg.Done()
	}(c)

	fmt.Printf("Message: %v", <-c)

	wg.Wait() // Wait until concurrent call is done

	// Output:

}

func Example_waitForOneFastest() {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1) // tells wg to consider only one concurrent process -- i.e. shortest
	go func(ch chan<- string) {
		time.Sleep(time.Millisecond * 500)
		ch <- "Hello"
		wg.Done()
	}(c)

	go func(ch chan<- string) {
		time.Sleep(time.Millisecond * 10)
		ch <- "World"
		wg.Done()
	}(c)

	fmt.Printf("Message: %v", <-c)

	wg.Wait() // Wait until concurrent call is done

	// Output:
	// Message: World
}

func Example_waitForTwo() {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2) // tells wg to consider only one concurrent process -- i.e. shortest
	go func(ch chan<- string) {
		time.Sleep(time.Millisecond * 500)
		ch <- "Hello"
		wg.Done()
	}(c)

	go func(ch chan<- string) {
		time.Sleep(time.Millisecond * 10)
		ch <- "World"
		wg.Done()
	}(c)

	fmt.Printf("Message 1: %v\n", <-c)
	fmt.Printf("Message 2: %v", <-c)

	wg.Wait() // Wait until concurrent call is done

	// Output:
	// Message 1: World
	// Message 2: Hello
}
