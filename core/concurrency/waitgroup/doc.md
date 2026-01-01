# Waitgroup

The `sync.WaitGroup` provides a more reliable way to manage goroutines than arbitrary `sleeps`. It is used to wait for a collection of goroutines to finish their execution.

## Pre-Go 1.25

This [example](./waitgroup/ex1/main.go) demonstrate the traditional implementation using `wg.Add()` and `wg.Done()`.

## Go 1.25+

This [example](./waitgroup/ex2/main.go) demonstrates the simpler, modern approach using the `wg.Go()` method, available from Go 1.25 onwards.