# Context

A context type represents signals to denotes deadlines, cancellation and scope values.

## References

* [Matt Kodvb - Go Class: 25 Context](https://www.youtube.com/watch?v=0x_oUlxzw5A&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=25)

## Example 1 

This example demonstrate the use of context with cancellation.

```go
ctx := context.Background()
ctx, cancel := context.WithCancel(ctx)

go func() {
	log.Println("Waiting for cancellation ...")
	<-ctx.Done() // Blocks routine until cancelled called
	log.Fatalf("Cancelled")
}()

for {
	time.Sleep(10 * time.Second)
	cancel()
}
```

Working example[./example/context/ex1/main.go](../example/context/ex1/main.go)

## Example 2

This example demonstrate the use of select to determine when cancel is called.

```go
var tm int
flag.IntVar(&tm, "time", 1, "time to sleep")
flag.Parse()

ctx, cancel := context.WithCancel(context.Background())
go func(cx context.Context) {
	select {
	case <-time.After(2 * time.Second): 
		fmt.Println("After 2 seconds") // This will print after two seconds
	case <-cx.Done():
		fmt.Println("Cancelled")
	}
}(ctx)

time.Sleep(time.Duration(tm) * time.Second)
cancel() // cancel called but if this called after 2 seconds this will not have any effect
time.Sleep(1 * time.Second) // This is to keep the goroutine
```

The [working example is here ./example/context/ex2/main.go](../example/context/ex2/main.go). To run the example:

* No flag - `go run main.go` this will sleep the main routine for 1 second
* With flag - `go run main.go -time=<seconds>` this will sleep the main routine for specified seconds

## Example 3

This example demonstrate a REST server catching a request from a browser. 

```go
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// This prints to STDOUT to show that processing has started
		fmt.Println("processing request")
		// We use `select` to execute a piece of code depending on which
		// channel receives a message first
		select {
		case <-time.After(2 * time.Second):
			// If we receive a message after 2 seconds
			// that means the request has been processed
			// We then write this as the response
			w.Write([]byte("request processed"))
		case <-ctx.Done():
			// If the request (e.g. browser closes) is cancelled in less than 2 secs
			fmt.Println("request cancelled")
		}
	}))
```

The [working example is here ](../example/context/ex3/main.go). To run this example, run this command `go run main.go`. Open a browser with the url: `localhost:8000`. You will see the message in your commandline `processing request`.

## Example 4

This example demonstrates context with timeout.

```go
ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()

log.Println("Process is blocked for 10 secs")
<-ctx.Done()
log.Println("Exiting...")
```

The [working example is here ./example/context/ex4/main.go](../example/context/ex4/main.go). To run this execute this example run this command `go run main.go`. The output could look like this:

```
2022/11/06 16:37:11 Process is blocked for 10 secs
2022/11/06 16:37:21 Exiting...
```

## Example 5

This example demonstrate the use of context to pass values.

```go
ctx := context.Background()

k := key("hello")
v := value("word")
ctx = context.WithValue(ctx, k, v)
fmt.Println(ctx.Value(k))
```

The [working example is here ./example/context/ex5/main.go](../example/context/ex5/main.go). Run the example through this command `go run main.go`.