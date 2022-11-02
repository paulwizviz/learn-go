# Concurrency

A concurrent program is made up of code blocks that can run at the same time (see official [document](https://www.golang-book.com/books/intro/10)).
However, concurrency is not the same as parallelism (see [Rob Pike](https://go.dev/blog/waza-talk))

NOTE: To execute your working example run `go run <path to the example>`

## Goroutines

A Go routine is a function with the keyword `go` added to a function. It enable you to create a function that can run concurrently with other functions.

<u>Example 1</u>

This example demonstrates a goroutine running within a main routine where the main ends before goroutine completes.

```go
func main() {
	go func() {
		fmt.Println("Hello") // This blocks the goroutine for a period of time
							 // io operations are typically blocked
	}()

    // The main routine executes immediately after goroutine
	// it ends before the goroutine ends killing the goroutine
}
```
[working example in ../example/concurrency/goroutine/ex1/main.go](../example/concurrency/goroutine/ex1/main.go)

<u>Example 2</u>

This example demonstrates the main routine sleeping for 100ms so the goroutine can complete its routine.

```go
func main() {
	go func() {
		fmt.Println("Hello, 世界")
	}()

	time.Sleep(100 * time.Microsecond) // Get the main routine to sleep for a period of time
	                                   // if the sleep period is longer than goroutine sleep
									   // you will see "Hello, 世界"
}
```
[Working example in ../example/concurrency/goroutine/ex2/main.go](../example/concurrency/goroutine/ex2/main.go)

<u>Example 3</u>

This example demonstrates the use of waitgroup to manage goroutines including tearing down. In this example there are 3 goroutines (main routine included).

```go
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Goroutine 1")
		wg.Done() // decrement count
	}()

	go func() {
		fmt.Println("Goroutine 2")
		wg.Done() // decrement count
	}()
	fmt.Println("Number of Goroutine before wait: ", runtime.NumGoroutine()) // This will report 3 goroutines (main is also a goroutine)
	wg.Wait() // Cause the main routine to pause until all routine is done
	fmt.Println("Number of Goroutine after wait: ", runtime.NumGoroutine()) // This will report 1 goroutine (main routine)
}
```
[Working example in ../example/concurrency/goroutine/ex3/main.go](../example/concurrency/goroutine/ex3/main.go)

<u>Example 4</u>

This examples demonstrates the use of mutex to synchronize shared variables across goroutines. Without mutex the shared variables might be updated in the wrong sequence. For example it might look like this:

```
Initial shared value: 0
 -- 
2 Before decrement: 0
1 Before increment: 0
2 After decrement: -1
 --- 
1 After increment: 0
 --- 
1 Before increment: 0
1 After increment: 1
 --- 
2 Before decrement: 1
2 After decrement: 0
 --- 
```

```go
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var mutext *sync.Mutex = &sync.Mutex{}

	var value = 0
	fmt.Printf("Initial shared value: %v\n -- \n", value)

	go func() {
		for i := 0; i < 2; i++ {
			time.Sleep(10 * time.Millisecond)
			func() {
				mutext.Lock() // If another gorountine holds the lock this will wait until holder release it hence blocking the completion of the goroutine
				fmt.Printf("1 Before increment: %v\n", value)
				value = value + 1
				fmt.Printf("1 After increment: %v\n --- \n", value)
				mutext.Unlock() 
			}()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 2; i++ {
			time.Sleep(10 * time.Millisecond)
			func() {
				mutext.Lock() // 
				fmt.Printf("2 Before decrement: %v\n", value)
				value = value - 1
				fmt.Printf("2 After decrement: %v\n --- \n", value)
				mutext.Unlock()
			}()
		}
		wg.Done()
	}()

	wg.Wait()
}
```

Mutex locking ensure that the shared variable is updated in the correct sequence:

```
Initial shared value: 0
 -- 
1 Before increment: 0
1 After increment: 1
 --- 
2 Before decrement: 1
2 After decrement: 0
 --- 
2 Before decrement: 0
2 After decrement: -1
 --- 
1 Before increment: -1
1 After increment: 0
 --- 
```


## Channels

A channel provide a way goroutines to communicate with each other. There are two types of channels:

* <b>Unbuffered channel</u> -- this type of channel requires a receiving channel to be available when a message is emitted from a sending channel.
* <b>Buffered channel</u> -- For this type of channel, sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

<u>Example 1</u>

This example demonstrates the operation of an unbuffered channel where the receiving channel is blocked until a message is emitted from the emitting channel.

```go
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
[Working example in ../example/concurrency/channel/ex1/main.go](../example/concurrency/channel/ex1/main.go)

<u>Example 2</u>

This example demonstrates a deadlock where there the receiving channel is expecting more messages than being sent.

```go
func main(){
	c := make(chan string)
	go func(ch chan<- string) {
		for i := 1; i < 6; i++ { // 
			ch <- "hello" 
		}
	}(c)
	for {
		fmt.Println(<-c) // After five incoming messagevs you get a deadlock and causing the main routine to panic
	}
}
```
[Working example in ../example/concurrency/channel/ex2/main.go](../example/concurrency/channel/ex2/main.go)

<u>Example 3</u>

This example demonstrates the closing of a channel before a message is sent.

```go
func main() {
	c := make(chan string)
	go func(ch chan<- string) {
		fmt.Println("Hello world") // 1 - This block the goroutine
		ch <- "Hello" // 3 - This is sent after channel is closed causing the routine to panic
	}(c)
	close(c) // 2 - Whilst the goroutine is block, this closes before the message hello is sent
	time.Sleep(1 * time.Millisecond)
}
```
[Working example in ../example/concurrency/channel/ex3/main.go](../example/concurrency/channel/ex3/main.go)

<u>Example 4</u>

This example demonstrates a technique to determine if a channel is opened or closed.

```go
func main() {
	c := make(chan string)
	go func(ch chan<- string) {
		defer close(ch) // Close channel at the end of the goroutine
		ch <- "Hello"
	}(c)

	v, opened := <-c // receive message before it is closed
	if opened {
		fmt.Printf("The channel is opened: %v\n", opened) // opened is true
		fmt.Printf("Message is: %v\n", v)
	}

	v, opened = <-c
	if !opened {
		fmt.Printf("The channel is opened: %v\n", opened) // opened is false
		fmt.Printf("Message is: %v\n", v)
	}
}
```
[Working example in ../example/concurrency/channel/ex4/main.go](../example/concurrency/channel/ex4/main.go)

<u>Example 5</u>

This example demonstrates the use of range to determined if a channel has been drained of all sending messages

```go
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
```
[Working example in ../example/concurrency/channel/ex5/main.go](../example/concurrency/channel/ex5/main.go)

<u>Example 6</u>

This example demonstrate the concept of a buffered channel.

```go
func main() {
	c := make(chan string, 2) // Create a channel with two slices of strings
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

	c <- "Hello"
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

	c <- "World"
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

	c <- "Hola" // Send a message to channel when its buffer is full
	            // this will cause a deadlock as there is no receiving
				// channel to drain the buffer.
				// fatal error: all goroutines are asleep - deadlock!
	fmt.Printf("Channel capacity: %v Length: %v\n", cap(c), len(c))

}
```
[Working example in ../example/concurrency/channel/ex6/main.go](../example/concurrency/channel/ex6/main.go)

<u>Example 7</u>

This example demonstrates the use of select to protect data injestion into a channel when the buffer is full and there is no receiving channel

```go
func main() {
	c := make(chan int, 2)
	go func(ch chan int) {
		defer close(c)
		for i := 0; i < 10; i++ {
			select {
			case ch <- i:
				fmt.Println("Pushing value: ", i)
			default:
				// This is called when the channel buffer is full.
				fmt.Printf("Don't push %v to channel\n", i)
			}
		}
	}(c)
	time.Sleep(1 * time.Millisecond)
}
```
[Working example in ../example/concurrency/channel/ex7/main.go](../example/concurrency/channel/ex7/main.go)

<u>Example 8</u>

This example demonstrates the use of select to determine the first goroutine to send a signal to the main routine.

```go
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func(ch chan string) {
		time.Sleep(10 * time.Millisecond)
		ch <- "Sleep for 10 ms"
	}(c1)

	go func(ch chan string) {
		time.Sleep(1 * time.Millisecond)
		ch <- "Sleep for 1 ms"
	}(c2)

	c3 := time.After(20 * time.Millisecond) // This function returns a channel

	for {
		select {
		case msg := <-c1: // This channel blocks the main routine for 10 ms
			fmt.Println(msg) // This prints next
		case msg := <-c2: // This channel blocks the main routine for 1 ms
			fmt.Println(msg) // This prints first
		case <-c3: // This channel blocks the main routine for 20 ms
			fmt.Println("Timeout") // This prints last
			return
		}
	}
}
```
[Working example in ../example/concurrency/channel/ex8/main.go](../example/concurrency/channel/ex8/main.go)

