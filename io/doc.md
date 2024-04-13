# IO Package

Go IO package provide interfaces and functionalities that wrap io operations and faciliate custom implementation of IO operations.

## Standard IO

Standard IO provides operation to read and write to shell.

Here is an example implementing [stdin](./stdin/main.go)

## Implementing IO readers and writers

Here are examples of demonstrating custom implementations of:

* [Reader](./reader/readers_test.go)
* [Writer](./writer/writers_test.go)