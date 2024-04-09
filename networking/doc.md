# Networking

This section illutrates the use standard library to build basic aspect of networking component.

For advanced networking patterns, please refer to [my other github repositories](https://github.com/paulwizviz/go-networking).

## TCP and Unix Socket programming

These examples demonstrate the use of standard library `net` package.

* [Auto assigning port](./auto-assign-port/main.go) - this demonstrates technique to auto assign a port.
* [Unix socket programming](./unix-socket/main.go) - this demonstrates client server via unix socket. To run this example, run the command:
```bash
go run networking/unix-socket/main.go -message="<your message here>"
```
This will send your message from a client to the server via `./temp/socket`.

## Http programming

These examples demonstrate the use of standard library `net/http` package and `httptest`.

* [Http client](./http-client/client_test.go) - this demonstrates http client programming.
* [Http server](./http-server/main.go) - this demonstrates http server programming.

## URL handling

These examples demonstrates the use of standard library `net/url` operations.

* [URL parsing](./urlparse/url_test.go) - This demonstrate operation to parse URL.
