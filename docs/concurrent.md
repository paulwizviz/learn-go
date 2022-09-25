# Concurrency

A concurrent program is made up of code blocks that can run at the same time (see official [document](https://www.golang-book.com/books/intro/10)).
However, concurrency is not the same as parallelism (see [Rob Pike](https://go.dev/blog/waza-talk))

## Goroutines

A Go routine is a function with the keyword `go` added to function. It enable you to create a function that can run concurrently with other functions.

<u>Scenario 1</u>

This demonstrates a goroutine running within a main routine where the main ends before goroutine completes.
```
package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Hello") // This will not print
	}()
}
```
[Run in playground](https://go.dev/play/p/suSAp8PWJdB)

<u>Scenario 2</u>

This demonstrate the main routine sleeping for 100ms so the goroutine can complete its routine.

```
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello, 世界")
	}()

	time.Sleep(100 * time.Millisecond)
}
```
[Run in playground](https://go.dev/play/p/mye4V2qkMXr)

## Channels

A channel provide a way for two goroutines or a main routine to communicate with each other.

<u>Scenario 1</u>

This scenario demonstrates a receiving channel blocking the main routine until a message is sent via the sending channel

```
   package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan string)

	go func(ch chan<- string) {
		time.Sleep(500 * time.Millisecond)
		ch <- "Hello world"
	}(c)

	timeBefore := time.Now()
	fmt.Println(timeBefore.Unix())
	msg := <-c // the main routine is blocked until it receive message from goroutine after 500ms
	timeAfter := time.Now()
	interval := time.Duration(500 * time.Millisecond)
	if interval < timeAfter.Sub(timeBefore) {
		fmt.Printf("Message: %s", msg) // This will print after 500ms
	}
}
```
[Working example](../example/concurrency/channel/ex1/main.go)