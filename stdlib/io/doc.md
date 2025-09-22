# IO Package

Go IO package provide interfaces and functionalities that wrap io operations and faciliate custom implementation of IO operations.

## Standard IO

Standard IO provides operation to read and write to shell.

Here is an example implementing [stdin](./stdin/main.go)

## Implementing IO readers and writers

Here are examples of demonstrating custom implementations of:

* [Reader](#io-reader)
* [Writer](#io-writer)

### IO Reader

[Example 1](./reader/stringreader_test.go) - This example derive a reader from standard package stringreader. We create a buffer of smaller than the initiate data. The read method is invoked and the reader store the index to the last string.

[Example 2](./reader/stringreader_test.go) - This example involves the use of a custom reader.

### IO Writer

[Writer](./writer/writers_test.go)
