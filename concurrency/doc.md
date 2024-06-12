# Concurrency

A concurrent program is made up of code blocks that can run at the same time (see official [document](https://www.golang-book.com/books/intro/10)). However, concurrency is not the same as parallelism (see [Rob Pike](https://go.dev/blog/waza-talk))

There are three basic elements of Go concurrency:

* Goroutine;
* Waitgroup;
* Channels.

# Goroutines

Adding the keyword `go` added to a function denotes a code block that can run in parallel if a processor supports it.

<u>Example 1</u>

This [working example](./goroutine/ex1/main.go) demonstrates a goroutine where main routine ends before goroutine completes.

<u>Example 2</u>

This [Working example](./goroutine/ex2/main.go) demonstrates the main routine sleeping for 1s whilsts two Go routines run concurrently.

## Channels

A channel is a mechanism for goroutine to communicate with each other. There are two types of channels:

* <b>Unbuffered channel</b> -- this type of channel requires a receiving channel to be available when a message is emitted from a sending channel.
* <b>Buffered channel</b>

<u>Example 1</U>

This [working example](./channel/ex1/main.go) demonstrates the operation of an unbuffered channel where the receiving channel is blocked until a message is emitted from the emitting channel.

<u>Example 2</u>

This [working example](./channel/ex2/main.go) demonstrates a deadlock where there the receiving channel is expecting more messages than being sent.

<u>Example 3</u>

This [working example](./channel/ex3/main.go) demonstrates the closing of a channel before a message is sent.

<u>Example 4</u>

This [working example](./channel/ex4/main.go) demonstrates a technique to determine if a channel is opened or closed.

<u>Example 5</u>

This [working example](./channel/ex5/main.go) demonstrates the use of range to determined if a channel has been drained of all sending messages.

<u>Example 6</u>

This [working example](./channel/ex6/main.go) demonstrate the concept of a buffered channel.

<u>Example 7</u>

This [working example](./channel/ex7/main.go) demonstrates the use of select to protect data injestion into a channel when the buffer is full and there is no receiving channel.

<u>Example 8</u>

This [working example](./channel/ex8/main.go) demonstrates the use of select to determine the first goroutine to send a signal to the main routine.

## Waitgroup

Waitgroups are elements that are used to manage goroutines which includes operations to wait until all goroutines are completed.

Please refer to these [working examples](./waitgroup/main.go).

## Concurrency patterns

Please refer to [GitHub repository for examples](https://github.com/paulwizviz/go-concurrency.git)

## References

* Matt Kodvb
    * [Go Class: 22 What is Concurrency?](https://www.youtube.com/watch?v=A3R-4ZYBqvE&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
    * [Go Class: 23 CSP, Goroutines, and Channels](https://www.youtube.com/watch?v=zJd7Dvg3XCk&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
    * [Go Class: 24 Select](https://www.youtube.com/watch?v=tG7gII0Ax0Q&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
    * [Go Class: 26 Channels in Detail](https://www.youtube.com/watch?v=fCkxKGd6CVQ&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)