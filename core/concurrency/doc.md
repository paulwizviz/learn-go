# Concurrency

Concurrency is one of Go's most powerful and discussed features. My initial exploration has been to understand the distinction between concurrency and parallelism, a concept Rob Pike explains well (see official [document](https://www.golang-book.com/books/intro/10) and [Rob Pike's talk](https://go.dev/blog/waza-talk)). My notes here document my observations on the core components of Go's concurrency model.

## Topics

* [Channels](./channel/doc.md)
* [Goroutines](./goroutine/doc.md)
* [Waitgroup](./waitgroup/doc.md)
* [Concurrency patterns](https://github.com/paulwizviz/advanced-go/blob/main/docs/concurrent.md)

## References

* Matt Kodvb
  * [Go Class: 22 What is Concurrency?](https://www.youtube.com/watch?v=A3R-4ZYBqvE&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
  * [Go Class: 23 CSP, Goroutines, and Channels](https://www.youtube.com/watch?v=zJd7Dvg3XCk&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
  * [Go Class: 24 Select](https://www.youtube.com/watch?v=tG7gII0Ax0Q&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)
  * [Go Class: 26 Channels in Detail](https://www.youtube.com/watch?v=fCkxKGd6CVQ&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6)