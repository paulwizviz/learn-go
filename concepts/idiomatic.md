# Idiomatic Go

My research into Go has led me to focus on the concept of "idiomatic" coding.
This refers to a style of programming that is natural and conventional within a
specific language. Essentially, idiomatic Go is about writing code that looks and feels like it belongs in the Go ecosystem—code that's straightforward, performant, and easy for other Gophers to read and contribute to.

Here are some of my observations.

## Simplicity

Idiomatic Go embraces this by avoiding overly complex designs, deep inheritance hierarchies, or heavy use of generics. The code should be easy to read and understand without a lot of mental overhead.

## Explicit Error Handling

Go doesn't use exceptions. Idiomatic Go code explicitly checks for errors using the if err != nil pattern. This makes error paths clear and forces the developer to handle them gracefully.

## Concurrency

Go's built-in concurrency model (goroutines and channels) is central to the language. Idiomatic Go uses these features to write concurrent programs that are easy to reason about and free from data races. The philosophy is "Don't communicate by sharing memory; share memory by communicating."

Refer to my [concurrency programming example](../concurrency/doc.md).

## Naming Conventions

Idiomatic Go uses `CamelCase` for exported identifiers (public) and `camelCase` for unexported ones (private).

A distinctive feature of idiomatic Go is the use of short variable names. This is a deliberate choice guided by a simple but powerful principle: **the length of a variable's name should be proportional to its scope.**

* **Small Scope:** For a variable used only within a few lines (like a loop index), a short name like `i`, `k`, or `r` is ideal. The context is immediate, so a longer name would be unnecessary visual noise.

* **Larger Scope:** The farther a variable is used from its declaration, the more descriptive its name should be. A variable used throughout a function might be `count` or `customerID`, while a package-level variable would have a fully descriptive name like `MaxActiveConnections`.

* **Context is Key:** Go's emphasis on small, focused packages means the surrounding code provides strong context. A variable `u` in a `users` package is clearly a user. This avoids redundant names like `userObject`.

In essence, the convention encourages names that are concise but not cryptic, balancing brevity with clarity.

## Package Layout

The file system structure defines packages. Idiomatic Go keeps packages small focused, and well-named. A package should do one thing and do it well. Refer to my observations of [Go Project Structure](./layout.md).

## Interfaces

Idiomatic Go favours composition over inheritance and uses small, single-method interfaces. Instead of a class inheriting behavior, a Go struct implements an interface by simply having the required method. This promotes loose coupling and flexible design.

Refer to [interfaces working example](../customs/doc.md).

## References

* [Blog - Idiomatic Go](https://dmitri.shuralyov.com/idiomatic-go)
* [Effective Go](https://go.dev/doc/effective_go) - "Official"
* [Idiomatic Go Naming Conventions (Golang)](https://www.youtube.com/watch?v=yQUAHpEvb9A)
* Mario Carrion
  * [5 Tips for Writing Idiomatic Code in Golang - Part 1](https://www.youtube.com/watch?v=TYQH4Rc6hwQ)
  * [5 Tips for Writing Idiomatic Code in Golang - Part 2](https://www.youtube.com/watch?v=lfQ4qLcE3Bo)
  * [5 Tips for Writing Idiomatic Code in Golang - Part 3](https://www.youtube.com/watch?v=qCg-FIkcJZw)
* [Package naming recommendation](https://go.dev/blog/package-names)
* [Rob Pike's Go Proverbs](http://go-proverbs.github.io/)
* [Using Go modules](https://go.dev/blog/using-go-modules)
