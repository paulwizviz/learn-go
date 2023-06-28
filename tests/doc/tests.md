# Tools

This section discusses Go tools for techniques for testing.

## Unit Testing

<u>Basic testing scenario:</u>

The idiomatic approach approach to unit testing is to use slices of scenarios and asserting using method provided by the testing package.

Here is a [working example](../examples/unit/ex1/ex1_test.go)

<u>Using test main</u>

In situation where you need a complex setup where you have tests spread across multiple files and you need a common setup use [TestMain](https://medium.com/goingogo/why-use-testmain-for-testing-in-go-dafb52b406bc).

Here is a [working example](../examples/unit/ex2/). Go into the package, and run the command `go test -v ./...`


## Benchmarking

The following are useful resources to learn about Go benchmarking.

* [Benchmarking in Golang: Improving function performance](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)
* [The basics of benchmarking Go applications | Smart Go](https://www.youtube.com/watch?v=pF7hMugCLZU)

<u>Basic setup</u>

This [example](../examples/benchmark/ex1/oneoff_test.go), demonstrate the most basic setup.

<u>Scenario based setup</u>

This [example](../examples/benchmark/ex2/scenarios_test.go), demonstrate a setup based on multiple scenarios.

## Fuzzing

In programming and software development, fuzzing or fuzz testing is an automated software testing technique that involves providing invalid, unexpected, or random data as inputs to a computer program[wiki](https://en.wikipedia.org/wiki/Fuzzing).

Here are some useful resources:

* [GopherCon 2022: Katie Hockman - Fuzz Testing Made Easy](https://www.youtube.com/watch?v=7KWPiRq3ZYI)

<u>Identifying panic</u>

This is one use case where Fuzz test is used to determine when a function panic. A working example is [here](../examples/fuzz/ex1/).

