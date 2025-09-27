# Concurrency

Concurrency is one of Go's most powerful and discussed features. My initial exploration has been to understand the distinction between concurrency and parallelism, a concept Rob Pike explains well (see official [document](https://www.golang-book.com/books/intro/10) and [Rob Pike's talk](https://go.dev/blog/waza-talk)). My notes here document my observations on the core components of Go's concurrency model.

* [Channels](#channels)
* [Goroutines](#goroutines)
* [Waitgroup](#waitgroup)
* [Concurrency patterns](#concurrency-patterns)
* [References](#references)

## Channels

Please refer to [channel](./channel/doc.md) for details

## Goroutines

Adding the `go` keyword before a function call creates a goroutine, which is a lightweight thread of execution.

Here are my observations from the following examples:

* [Example 1](./goroutine/ex1/main.go) was a clear demonstration for me of how a program's main thread will not wait for a goroutine to complete. The program simply exits.
* [Example 2](./goroutine/ex2/main.go) shows how a simple, though impractical, `sleep` can give concurrent routines time to run. This highlighted for me the need for a more robust mechanism to manage the lifecycle of goroutines.

## Waitgroup

The `sync.WaitGroup` provides a more reliable way to manage goroutines than arbitrary `sleeps`. My notes in this [working example](./waitgroup/main.go) show how to use it to wait for a collection of goroutines to finish their execution before the main program exits.

## Concurrency patterns

Please refer to [Advance Go](https://github.com/paulwizviz/advanced-go.git)

## References

* Matt Kodvb
  * [Go Class: 22 What is Concurrency?](https://www.youtube.com/watch?v=A3R-4ZYBqvE&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
  * [Go Class: 23 CSP, Goroutines, and Channels](https://www.youtube.com/watch?v=zJd7Dvg3XCk&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
  * [Go Class: 24 Select](https://www.youtube.com/watch?v=tG7gII0Ax0Q&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
  * [Go Class: 26 Channels in Detail](https://www.youtube.com/watch?v=fCkxKGd6CVQ&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)