# What is Idiomatic Go?

Idiomatic Go refers to the set of conventions, best practices, and common patterns that Go programmers use to write clear, simple, and efficient code. It's about following the "Go way" of doing things, prioritizing readability and maintainability over clever or complex solutions.

## Key Principles of Idiomatic Go

* **Simplicity:** Go's language features are minimalist. Idiomatic Go embraces this by avoiding overly complex designs, deep inheritance hierarchies, or heavy use of generics. The code should be easy to read and understand without a lot of mental overhead.

* **Explicit Error Handling:** Go doesn't use exceptions. Idiomatic Go code explicitly checks for errors using the if err != nil pattern. This makes error paths clear and forces the developer to handle them gracefully.

* **Concurrency:** Go's built-in concurrency model (goroutines and channels) is central to the language. Idiomatic Go uses these features to write concurrent programs that are easy to reason about and free from data races. The philosophy is "Don't communicate by sharing memory; share memory by communicating."

* **Naming Conventions:** Naming is crucial for readability. Idiomatic Go uses CamelCase for exported identifiers (public) and snake_case for unexported ones (private). Variable names are often short and descriptive (e.g., i for a loop index, r for a reader).

* **Package Layout:** The file system structure defines packages. Idiomatic Go keeps packages small, focused, and well-named. A package should do one thing and do it well.

* **Interfaces:** Idiomatic Go favors composition over inheritance and uses small, single-method interfaces. Instead of a class inheriting behavior, a Go struct implements an interface by simply having the required method. This promotes loose coupling and flexible design.

Essentially, idiomatic Go is about writing code that looks and feels like it belongs in the Go ecosystem—code that's straightforward, performant, and easy for other Gophers to read and contribute to.

## Go Modules and Packages

In Go, the two primary components for organising code are **modules** and **packages**. Understanding their relationship is fundamental to structuring a project idiomatically.

* **Module:** A module is a collection of related Go packages that are versioned together as a single unit. It is defined by a `go.mod` file at the root of the project directory. This file declares the module's path (its unique name), the version of Go it is built with, and its dependencies on other modules. A repository typically contains a single module.

* **Package:** A package is a directory containing one or more Go source files. All files within a directory must belong to the same package, declared using the `package` keyword at the top of each file. A package provides a distinct namespace for its exported identifiers (those starting with a capital letter) and represents a single, coherent unit of functionality.

Consider the following project structure:

```sh
/my-project
├── go.mod         // Defines the 'my-project' module
├── main.go        // package main
|
└───/calculator/
    ├── add.go     // package calculator
    └── subtract.go// package calculator
```

In this example:

1. The entire `my-project` directory constitutes a **module**, as defined by the `go.mod` file.
2. The `main.go` file belongs to the special `main` **package**, which signifies an executable program.
3. The `calculator` directory forms a separate **package** named `calculator`. It contains related functions for performing calculations, which can be imported and used by other packages (like `main`).

## The `cmd` and `internal` Directories

Go's official documentation and the broader community recommend a specific directory layout to solve common problems:

* `cmd/`: This directory should contain the source code for each of your main applications or executables. For example, if your project has a web server and a background worker, they would live in cmd/web and cmd/worker, respectively. This clearly separates your application entry points (which contain main packages) from the reusable library code.

* `internal/`: This is a special directory used for private packages that cannot be imported by other repositories. It's the primary mechanism for enforcing encapsulation. Any code you don’t want to be part of your public API or shared with other modules should go here.

It is also important to note what to avoid: 

* `src/` directory. This is a common pattern in other languages but is considered an anti-pattern in Go, as the entire Go workspace is already a source-code directory.

* `pkg/` folder is often unnecessary and can become a dumping ground for packages.

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