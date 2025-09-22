# Context

We learn a context type represents signals to denotes deadlines, cancellation and scope values.

## Working Examples

### Example 1

This [working example](./ex1/main.go) demonstrate the use of context with cancellation.

### Example 2

This [Working example](./ex2/main.go). demonstrate the use of select to determine when cancel is called.

To run the example:

* No flag - `go run main.go` this will sleep the main routine for 1 second
* With flag - `go run main.go -time=<seconds>` this will sleep the main routine for specified seconds

### Example 3

This [Working example](./ex3/main.go) demonstrate a REST server catching a request from a browser.

To run the example, run the command `go run main.go`. Open a browser with the url: `localhost:8000`. You will see the message in your commandline `processing request`.

### Example 4

This [Working example](./ex4/main.go) demonstrates context with timeout.

To run this execute this example run this command `go run main.go`. The output could look like this:

```sh
2022/11/06 16:37:11 Process is blocked for 10 secs
2022/11/06 16:37:21 Exiting...
```

### Example 5

This [Working example](./ex5/main.go) demonstrate the use of context to pass values.

Run the example through this command `go run main.go`.

## References

* [Matt Kodvb - Go Class: 25 Context](https://www.youtube.com/watch?v=0x_oUlxzw5A&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=25)