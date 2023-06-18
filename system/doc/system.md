# System

This section discuss techniques for system programming.

## Processing incoming signals

Go [signal package](https://pkg.go.dev/os/signal) provides operations to handle incoming signals. A SIGHUP, SIGINT, or SIGTERM signal causes the program to exit.

<u>Executable examples</u>

[Source code](../examples/sig/main.go)

Note this is only applicable to Unix/macOS app

To see this example in action, follow these steps:

* STEP 1: Build app, run the command `go build -o $PWD/tmp/sig $PWD/example/system/sig`
* STEP 2: Run the app in `./temp/sig`
* STEP 3: Kill sig process or execute ctl-c
