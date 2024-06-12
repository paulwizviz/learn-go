# Concurrency

A concurrent program is made up of code blocks that can run at the same time (see official [document](https://www.golang-book.com/books/intro/10)). However, concurrency is not the same as parallelism (see [Rob Pike](https://go.dev/blog/waza-talk))

There are three basic elements of Go concurrency:

* Goroutine;
* Waitgroup;
* Channels.

## Goroutines

Adding the keyword `go` added to a function denotes a code block that can run in parallel if a processor supports it.

<u>Example 1</u>

This [working example](./goroutine/ex1/main.go) demonstrates a goroutine where main routine ends before goroutine completes.

<u>Example 2</u>

This [Working example](./goroutine/ex2/main.go) demonstrates the main routine sleeping for 1s whilsts two Go routines run concurrently.

## Channels

A channel is a mechanism for goroutines and the main routine to communicate with each other. There are two types of channels:

* <b>Unbuffered channel</b> -- This type of channel requires a receiving channel to be available when a message is emitted from a sending channel. If there is no receiver it will cause a deadlock.
* <b>Buffered channel</b> -- A buffer channel may receive message without any receiver. If you overflow the buffer it will cause a deadlock. If the receiver tries to access more message than is available in the buffer it will cause a deadlock.

<u>Example 1</U>

The following examples are based on unbuffered channel.

* [Example 1a](./channel/ex1a/main.go) we create an unbuffered channel without a receiving channel, causing Goroutine deadlock.
* [Example 1b](./channel/ex1b/main.go) demonstrates blocking operations with channel
* [Example 1c](./channel/ex1c/main.go) demonstrates a scenario with safeguard to check if channel is closed.
* [Example 1d](./channel/ex1d/main.go) demonstrates a scenario with no safeguard.
* [Example 1e](./channel/ex1e/main.go) demonstrates closing of channel prematurely.
* [Example 1f](./channel/ex1f/main.go) using alternative ways to signal closure of channel.
* [Example 1g](./channel/ex1g/main.go) using ranging to avoid goroutine deadlock

<u>Example 2</u>

The following examples are based on buffered channel.

* [Example 2a](./channel/ex2a/main.go) demonstrates a deadlock where there the receiving channel is expecting more messages than being sent.
* [Example 2b](./channel/ex2b/main.go) demonstrate the use of select to avoid sending data to buffered channel when it is full.


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