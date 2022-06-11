# Networking

This section discusses basic network programming techniques.

## Auto-assigning-port

* Executable examples
    * [auto-assigning-port](../example/network/auto-assign-port/)

## Proxying

There are two kinds of network proxies.

![Forward proxy](./img/forward-proxy-flow.svg)
![Reverse proxy](./img/reverse-proxy-flow.svg)

* Source
    * [What is a reverse proxy?](https://www.cloudflare.com/en-gb/learning/cdn/glossary/reverse-proxy/)

* Executable examples
    
* [Simple reverse proxy server](../example/network/proxy/).
  * STEP 1: Run the command `go run example/network/proxy/actual/main.go` to start the actual backend server (port 3031).
  * STEP 2: Run the command `go run example/network/proxy/middle/main.go` to start the proxy server (port 3030).
  * STEP 3: Use a client (curl, browser, etc) and target the proxy server (http://localhost:3030).
