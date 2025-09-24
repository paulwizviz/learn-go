# Concurrency

A concurrent program is made up of code blocks that can run at the same time (see official [document](https://www.golang-book.com/books/intro/10)). However, concurrency is not the same as parallelism (see [Rob Pike](https://go.dev/blog/waza-talk))

There are three basic elements of Go concurrency:

- [Concurrency](#concurrency)
  - [Channels](#channels)
  - [Goroutines](#goroutines)
  - [Waitgroup](#waitgroup)
  - [Concurrency patterns](#concurrency-patterns)
  - [References](#references)

## Channels

Please refer to [this section](./channel/doc.md) for details

## Goroutines

Adding the keyword `go` added to a function denotes a code block that can run in parallel if a processor supports it.

Here are the following examples:

* [Example 1](./goroutine/ex1/main.go) - demonstrates a goroutine where main routine ends before goroutine completes.
* [Example 2](./goroutine/ex2/main.go) demonstrates the main routine sleeping for 1s whilsts two Go routines run concurrently.

## Waitgroup

Waitgroups are elements that are used to manage goroutines which includes operations to wait until all goroutines are completed.

Please refer to these [working examples](./waitgroup/main.go).

## Concurrency patterns

Please refer to [GitHub repository for examples](https://github.com/paulwizviz/go-concurrency.git)

## References

- Matt Kodvb
  - [Go Class: 22 What is Concurrency?](https://www.youtube.com/watch?v=A3R-4ZYBqvE&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
  - [Go Class: 23 CSP, Goroutines, and Channels](https://www.youtube.com/watch?v=zJd7Dvg3XCk&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
  - [Go Class: 24 Select](https://www.youtube.com/watch?v=tG7gII0Ax0Q&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
  - [Go Class: 26 Channels in Detail](https://www.youtube.com/watch?v=fCkxKGd6CVQ&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)