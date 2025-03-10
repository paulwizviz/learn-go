# OS package

This package provide interfaces and functionalities to interact with, principally Unix-based, operating system.

## File

Example file operations [here](./file/file_test.go)

## Signals

Go [signal package](https://pkg.go.dev/os/signal) provides operations to handle incoming signals. A SIGHUP, SIGINT, or SIGTERM signal causes the program to exit.

Here is an [example source code](./sig/main.go). Note this is only applicable to Unix/macOS app

To see this example in action, follow these steps:

* STEP 1: Build app, run the command `go build -o $PWD/temp/sig $PWD/system/sig`
* STEP 2: Run the app in `./temp/sig`
* STEP 3: Kill sig process or execute ctl-c

## Shell

A shell is an program that enable users to interact with operating system's services.

Here are examples demonstrating ways to write apps to call shell-like operations:

* [Arguments and flags](./args/main.go) -- this demonstrate operations for apps to read command line aguments and, association with `flag` packages.
* [Shell execution](./shell/exec_test.go) -- this demonstrates operations to execute external applications from shell processes.
