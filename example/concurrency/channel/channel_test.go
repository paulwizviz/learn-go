package channel_test

import (
	"fmt"
	"time"
)

// This example demonstrate a receiving channel blocking
// the main process until a sending channel has sent its
// message.
func Example_blockingChannel() {
	c := make(chan string)
	go func(thing string, c chan string) {
		// This will sleep for 500 millisecond before
		// sending message
		time.Sleep(time.Millisecond * 500)
		c <- thing
	}("sheep", c)

	timeBefore := time.Now()
	msg := <-c // process is blocked until it receive message from channel after 500ms

	timeAfter := time.Now()
	fmt.Println(timeBefore.Unix())
	fmt.Println(timeAfter.Unix())
	interval := time.Duration(500 * time.Millisecond)
	if interval < timeAfter.Sub(timeBefore) {
		fmt.Printf("Message: %s", msg)
	}

	// Output:
	// Message: sheep
}

// This example shows go routine sending
// five messages to a sending channel
// and the main process expecting
// more than five messages
func Example_deadlock() {
	c := make(chan string)
	go func(ch chan<- string) {
		for i := 1; i < 6; i++ {
			ch <- "hello"
		}
	}(c)
	for {
		fmt.Println(<-c)
	}

	// You will not see this in this test framework. Expected output if run as part of a main process.
	// hello
	// hello
	// hello
	// hello
	// hello
	// fatal error: all goroutines are asleep - deadlock!

}

// This example shows the effect of closing receiving
// channel before any message is received
func Example_closeChanBeforeSending() {
	c := make(chan string)
	go func(ch chan<- string) {
		fmt.Println("Hello world") // This will not print
		ch <- "Hello"
	}(c)
	close(c)

	// You will not see output in this test framework. Expected output if run as part of a main process.
	// Hello world
	// panic: send on closed channel
}

// This example demonstrate a way of determining if
// a channel is closed
func Example_verifyClosedChannel() {
	c := make(chan string)
	go func(ch chan<- string) {
		defer close(ch)
		ch <- "Hello"
	}(c)

	v, ok := <-c
	if ok {
		fmt.Printf("The channel is opened: %v\n", ok)
		fmt.Printf("Message is: %v\n", v)
	}

	v, ok = <-c
	if !ok {
		fmt.Printf("The channel is opened: %v\n", ok)
		fmt.Printf("Message is: %v\n", v)
	}

	// Output:
	//
	// The channel is opened: true
	// Message is: Hello
	// The channel is opened: false
	// Message is:
}

// Range over a channel to ensure that the
// number of receiving messages matches
// the number sending messages
func Example_rangeChans() {
	c := make(chan string)
	go func(ch chan<- string) {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- "Hello"
		}
	}(c)

	for v := range c {
		fmt.Println(v)
	}

	// Output:
	// Hello
	// Hello
	// Hello
}

// Buffered channel sized for 2 messages
func Example_bufferedChan() {
	c := make(chan string, 2) // Set channel to accept to messages

	go func(ch chan<- string) {
		ch <- "Hello"
		ch <- "world"
	}(c)

	fmt.Println(<-c)
	fmt.Println(<-c)

	// Output:
	// Hello
	// world
}

// This example demonstrate a buffered channel
func Example_bufferedCha() {

	c := make(chan string, 2) // Create a channel with two slices of strings
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

	c <- "Hello"
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

	c <- "World"
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

	// // Push if only capacity exists
	ola := "Ola"
	select {
	case c <- ola:
		// This will not happen unless you make(chan string, 3)
	default:
		// Do nothing if the capacity of c exceeded
	}

	fmt.Printf("Message 1: %v\n", <-c)
	fmt.Printf("Message 2: %v\n", <-c)
	// fmt.Printf("Message 3: %v\n", <-c) // <-c is Ola when you make(chan string, 3)

	// Output:
	// Channel capacity: 2 Length: 0
	// Channel capacity: 2 Length: 1
	// Channel capacity: 2 Length: 2
	// Message 1: Hello
	// Message 2: World
}

// This demonstrate receivers without
// switching
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
		fmt.Println(<-c1) // This is block until the sender wakes up to send
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

// This example demonstrate the use of select
// to switch between multiple receiving channels
func Example_select() {
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

	// Output:
	// Every 10 milliseconds
	// Every 500 milliseconds
	// Every 10 milliseconds
	// Every 10 milliseconds
	// Every 10 milliseconds
	// Every 10 milliseconds
	// Every 500 milliseconds
	// Every 500 milliseconds
	// Every 500 milliseconds
	// Every 500 milliseconds

}
