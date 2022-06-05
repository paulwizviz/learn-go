package channel_test

import (
	"fmt"
	"time"
)

func count(thing string, c chan string) {
	for i := 1; i < 5; i++ {
		// This will sleep for 500 millisecond before
		// sending message
		time.Sleep(time.Millisecond * 500)
		c <- thing
	}
}

func Example_receiveOneMessage() {
	c := make(chan string)
	go count("sheep", c)
	timeBefore := time.Now()
	msg := <-c
	timeAfter := time.Now()
	fmt.Printf("Time interval: %v Message: %s", timeAfter.Sub(timeBefore), msg)
}

// This example shows routine sending
// six messages to a channel and rountine will
// report a deadlock error when receiving channel
// is expecting more than six messages
func Example_matchingSender() {
	c := make(chan string)
	go func(ch chan<- string) {
		for i := 1; i < 6; i++ {
			ch <- "hello"
		}
	}(c)
	for {
		fmt.Println(<-c)
	}

}

// This example shows the effect of closing receiving
// channel before any message is received
func Example_closeChanBeforeSending() {
	c := make(chan string)
	go func(ch chan<- string) {
		fmt.Println("Hello") // This will not print
		ch <- "Hello"
	}(c)
	close(c)

	// Output:
}

func Example_verifyClosedChannel() {
	c := make(chan string)
	go func(ch chan<- string) {
		defer close(ch)
		ch <- "Hello"
	}(c)

	v, ok := <-c
	if ok {
		fmt.Println("This is expected")
		fmt.Printf("Message is: %v\n", v)
	}

	v, ok = <-c
	if !ok {
		fmt.Printf("Channel is already closed. Value: %v", v)
	}

	// Output:
	// This is expected
	// Message is: Hello
	// Channel is already closed. Value:
}

func Example_bufferedChannels() {

	c := make(chan string, 2)
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

	c <- "Hello"
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

	c <- "World"
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

	// Push if only capacity exists
	select {
	case c <- "Ola":
		// This will not happen
	default:
		// Do nothing
	}

	fmt.Printf("Message 1: %v\n", <-c)
	fmt.Printf("Message 2: %v\n", <-c)

	// Output:
	// Channel capacity: 0 Length: 2
	// Channel capacity: 1 Length: 2
	// Channel capacity: 2 Length: 2
	// Message 1: Hello
	// Message 2: World
}

func Example_noSwitch() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			c1 <- "Every 500 milliseconds"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			c2 <- "Every 10 milliseconds"
			time.Sleep(time.Millisecond * 10)
		}
	}()

	// Channel is blocked by the slowest
	// so message will print the slowest
	// first then the next available one
	for i := 0; i < 5; i++ {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}

	// Output:
	// Every 500 milliseconds
	// Every 10 milliseconds
	// Every 500 milliseconds
	// Every 10 milliseconds
	// Every 500 milliseconds
	// Every 10 milliseconds
	// Every 500 milliseconds
	// Every 10 milliseconds
	// Every 500 milliseconds
	// Every 10 milliseconds

}

func Example_switch() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			c1 <- "Every 500 milliseconds"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			c2 <- "Every 10 milliseconds"
			time.Sleep(time.Millisecond * 10)
		}
	}()

	for i := 0; i < 10; i++ {
		// This will pull the very first sender
		// when it is available.
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
