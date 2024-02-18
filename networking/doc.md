# Networking

This section illutrates the use standard library to build aspect of networking component.

## Working Examples

<u>TCP and Unix Socket programming</u>

These examples demonstrate the use of standard library `net` package.

* [Auto assigning port](./auto-assign-port/main.go) - this demonstrates technique to auto assign a port.
* [Unix socket programming](./unix-socket/) - this demonstrates client server via unix socket. To run this example, open a shell and run the server `go run ./networking/unix-socket/server/main.go`. In another shell run the client `go run ./networking/unix-socket/client/main.go`.

<u>Http programming</u>

These examples demonstrate the use of standard library `net/http` package.

* [Http client](./http-client/client_test.go) - this demonstrates http client programming.
* [Http server](./http-server/main.go) - this demonstrates http server programming.
* [Webserver](./webserver) - these examples demonstrate web and file servers programming.

<u>URL handling</u>

These examples demonstrates the use of standard library `net/url` operations.

* [URL parsing](./urlparse/url_test.go) - This demonstrate operation to parse URL.

## References

* My collection of [advanced networking patterns](https://github.com/paulwizviz/go-networking) - This include the use of non-standard packages.