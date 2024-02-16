# Networking

This section illutrates the use standard library to build aspect of networking component.

For advance example, please refer to [github.com/paulwizviz/go-networking](https://github.com/paulwizviz/go-networking)

## TCP and Unix Socket programming

These examples demonstrate the use of standard library `net` package.

* [Auto assigning port](./auto-assign-port/main.go) - this demonstrates technique to auto assign a port.
* [Unix socket programming](./unix-socket/) - this demonstrates client server via unix socket. To run this example, open a shell and run the server `go run ./networking/unix-socket/server/main.go`. In another shell run the client `go run ./networking/unix-socket/client/main.go`.

## Http programming

These examples demonstrate the use of standard library `net/http` package.

* [Http client](./http-client/client_test.go) - this demonstrates http client programming.
* [Http server](./http-server/main.go) - this demonstrates http server programming.
* [Webserver](./webserver) - these examples demonstrate web and file servers programming.

## URL handling

These examples demonstrates the use of standard library `net/url` operations.

* [URL parsing](./urlparse/url_test.go) - This demonstrate operation to parse URL.