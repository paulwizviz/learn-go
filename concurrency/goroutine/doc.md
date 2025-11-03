# Goroutines

Adding the `go` keyword before a function call creates a goroutine, which is a lightweight thread of execution.

Here are my observations from the following examples:

* [Example 1](./goroutine/ex1/main.go) was a clear demonstration for me of how a program's main thread will not wait for a goroutine to complete. The program simply exits.
* [Example 2](./goroutine/ex2/main.go) shows how a simple, though impractical, `sleep` can give concurrent routines time to run. This highlighted for me the need for a more robust mechanism to manage the lifecycle of goroutines.