# Unix Socket Programming

This section describe techniques for Unix socket programming.

Here is a [working example](./main.go) demonstrating client server communications via unix socket.

To run this example, execute the command:

```bash
go run networking/unix-socket/main.go -message="<your message here>"
```

This will send your message from a client to the server via `./temp/socket`.