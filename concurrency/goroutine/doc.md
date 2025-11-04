# Goroutines

Adding the `go` keyword before a function call creates a goroutine, which is a lightweight thread of execution.

## Example 1

This [Example](./goroutine/ex1/main.go) demonstrates for me of how a program's main thread will not wait for a goroutine to complete. The program simply exits.

## Example 2

This [Example](./goroutine/ex2/main.go) shows how a simple, though impractical, `sleep` can give concurrent routines time to run. This highlighted for me the need for a more robust mechanism to manage the lifecycle of goroutines.