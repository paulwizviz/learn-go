# Unit Testing

* [Table Driven](#table-driven)
* [TestMain Setup](#testmain-setup)

## Table Driven

The idiomatic approach to writing unit test is to create test scenarios using slices. The scenarios should be separate from test runners. Here is how to structure your test:

* [Functions as subjects for testing](./ex1/ex1.go).
* [Test scenarios and runners](./ex1/ex1_test.go).

**Note**: Test scenarios could also be from a files.

## TestMain Setup

In situation where you need a common setup where there are multiple test runners and also spread across multiple test files. You can use [TestMain](https://medium.com/goingogo/why-use-testmain-for-testing-in-go-dafb52b406bc).

In our working example, we are simulating a common DB setup and test runners spread across multiple files.

* [The file with TestMain with first and second runners](./ex2/s_test.go)
* [The file with third runner](./ex2/z_test.go)
* [The file with fourth runner](./ex2/m_test.go)

Run the command `go test -v ./...` to execute the test.
