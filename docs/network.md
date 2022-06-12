# Networking

This section discusses basic network programming techniques.

## Auto-assigning-port

* Executable examples
    * [auto-assigning-port](../example/network/auto-assign-port/)

## Proxying

There are two kinds of network proxies.

![Forward proxy](./img/forward-proxy-flow.svg)
![Reverse proxy](./img/reverse-proxy-flow.svg)

### Reference materials
* [What is a reverse proxy?](https://www.cloudflare.com/en-gb/learning/cdn/glossary/reverse-proxy/)

### Executable examples

[Source code](../example/network/proxy/)

**Using httputil**

This example use httputil. To see the example in action, follow these steps:

  * STEP 1: Run the command `go run example/network/proxy/actual/main.go` to start the actual backend server (port 3031).
  * STEP 2: Run the command `go run example/network/proxy/httputil/main.go` to start the proxy server (port 3030).
  * STEP 3: Use a client (curl, browser, etc) and target the proxy server (http://localhost:3030).

**Using custom provxy**

This example use a custom proxy. To see the example in action, follow these steps:

  * STEP 1: Run the command `go run example/network/proxy/actual/main.go` to start the actual backend server (port 3031).
  * STEP 2: Run the command `go run example/network/proxy/custom/main.go` to start the proxy server (port 3030).
  * STEP 3: Use a client (curl, browser, etc) and target the proxy server (http://localhost:3030).

