# Design Patterns

## Worker Pool Patterns

A worker pool pattern using a combination of goroutine and channel.

This working example [./example/concurrency/worker/main.go](../example/patterns/worker/main.go) demonstrates the use of worker pool to build a fibonance sequence for a range of index using workers. To run the example, execute: `go run example/concurrency/worker/main.go -seq=<sequence of fibonance> -workers=<number of workers>`.

## Observer Patterns

This working example [./example/concurrency/pubsub/main.go](../example/patterns/pubsub/main.go) demonstrates an implementation of a pub sub (or listener) pattern using channels.