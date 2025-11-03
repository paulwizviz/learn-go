# Waitgroup

The `sync.WaitGroup` provides a more reliable way to manage goroutines than arbitrary `sleeps`. It is used to wait for a collection of goroutines to finish their execution.

My notes in these working examples demonstrate two approaches:

* **[Pre-Go 1.25](./waitgroup/ex1/main.go):** This example shows the traditional implementation using `wg.Add()` and `wg.Done()`.
* **[Go 1.25+](./waitgroup/ex2/main.go):** This example demonstrates the simpler, modern approach using the `wg.Go()` method, available from Go 1.25 onwards.