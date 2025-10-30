# Naming Conventions

Idiomatic Go uses `CamelCase` for exported identifiers (public) and `camelCase` for unexported ones (private).

A distinctive feature of idiomatic Go is the use of short variable names. This is a deliberate choice guided by a simple but powerful principle: **the length of a variable's name should be proportional to its scope.**

## Small Scope

For a variable used only within a few lines (like a loop index), a short name like `i`, `k`, or `r` is ideal. The context is immediate, so a longer name would be unnecessary visual noise.

## Larger Scope

The farther a variable is used from its declaration, the more descriptive its name should be. A variable used throughout a function might be `count` or `customerID`, while a package-level variable would have a fully descriptive name like `MaxActiveConnections`.

## Package Names

A package is a directory in a file system for grouping Go codes into related files. In the context of Go, a package is also a unit of abstraction.

A unit of abstraction is a concept or a representation of a entity that is defined by its essential characteristics while hiding its complex, unnecessary implementation details.

A package name should be short (no underscore or mixed case) and reflect the intended abstraction. If you wanted your package to represent an entity of interest to you, use name that best reflect accordingly. For example, you wanted to package to represent the characteristics (data types and functions) of a user of a system, name the package `user`. The content of the package should include data types or functions that represents the characteristics and behaviour of the entity.

If the name of an entity is long abbreviate it appropriately. For example, you wanted to represent a concept called job application, you could abbreviate it to `jobapp`. However, abbreviation should not be at the expense of clarity.

Avoid naming packages to reflect some collection such as `util` or `common`.

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