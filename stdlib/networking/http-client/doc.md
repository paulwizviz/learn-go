# HTTP Client

Here are techniques for creating HTTP Client.

Here are [working examples](./client_test.go).

## Topics

* [Draining Response Body](#draining-response-body)
* [Handling Body Closing Error](#handling-closing-error)

## Draining Response Body

You should drain the http.Response.Body before closing it primarily to **ensure that the underlying TCP connection can be reused** for subsequent HTTP requests. This is crucial for efficiency and preventing resource leaks.

Here's a more detailed explanation:

### How HTTP Keeps Connections Alive

For performance reasons, HTTP/1.1 and HTTP/2 often keep TCP connections alive after a request-response cycle is complete. This allows the client and server to send multiple requests and responses over the same persistent connection, reducing the overhead of establishing new connections for each request.

### Why Draining the Body Matters for Connection Reuse

* **Reading the Entire Body**: To reuse a persistent connection, both the client and the server need to have fully processed the current request and response. This includes reading the entire response body.
* **Unread Data Prevents Reuse**: If the client (in this case, your Go HTTP client) doesn't read the entire response body before closing it, the server might still be in the process of sending data. The unread data in the buffer can signal to the server that the client hasn't fully processed the response.
* **Connection Closure**: In such a scenario, the server might decide to close the TCP connection to avoid potential issues or to ensure proper protocol handling. This defeats the purpose of persistent connections and forces the client to establish a new connection for the next request, which is less efficient.
* **Resource Leaks (Less Common with Proper Close)**: While less likely with a defer resp.Body.Close(), if the body isn't fully read and the connection isn't properly closed, there could be lingering resources on the server side. However, the primary concern is usually connection reuse.

### When is Draining Particularly Important?

* **When you don't read the entire body**: If your application logic only needs to examine the headers or a small part of the body and doesn't consume the rest, you should definitely drain the remaining data.
* **For persistent connections**: If you intend to make subsequent requests to the same server over the same connection (which is the default behavior of http.Client), draining is crucial for efficient reuse.

### Summary Draining the Body

In summary, draining the http.Response.Body before closing it is a best practice to ensure proper handling of persistent connections, allowing for more efficient communication between your client and the server by enabling connection reuse. It signals that the entire response has been processed, preventing the server from prematurely closing the connection.

## Handling Closing Error

The error returned by resp.Body.Close() is an error type, just like any other error in Go. According to the io.Closer interface documentation, the Close method should return an error if any occurred during the closing process.

While closing an HTTP response body is usually a straightforward operation, there are potential (though less common) scenarios where it might return an error:

### Possible Errors from resp.Body.Close()

* **Underlying Connection Issues**: If the underlying network connection has already been closed or is in a bad state, the Close() operation might return an error related to network communication.
Resource Issues: In some rare cases, there might be issues with releasing resources associated with the body, although this is less common for HTTP responses compared to file operations.
* **Implementation-Specific Errors**: The specific implementation of the ReadCloser returned as the resp.Body could potentially have its own specific error conditions for closing.

### What You Can Do with the Error

Even though errors from resp.Body.Close() are not very frequent or critical in most HTTP client scenarios, it's still good practice to handle them, especially in long-running applications or critical systems.

Here's what you can do with the error:

* **Log the Error**: The most common approach is to log the error. This can be helpful for debugging if unexpected issues arise. You can use the log package or a more sophisticated logging library.

* **Increment Error Metrics**: If you are monitoring your application's health and performance, you might want to increment a counter for errors encountered during response body closing. This can help you track the frequency of such issues.

* **Conditional Retry (Potentially Risky)**: In very specific and controlled scenarios, you might consider a retry if closing the body fails due to a transient network issue. However, this should be done with caution, as blindly retrying a close operation might not always be safe or effective.

* **No Action (Ignoring the Error)**: In many simple HTTP client applications, the error from `resp.Body.Close()` is often ignored. The reasoning is that if closing the body fails, the underlying resources will eventually be reclaimed by the operating system when the program exits or the connection times out. However, explicitly ignoring errors can mask potential underlying problems, so logging is generally preferred.

### Summary Handling Closing Error

In summary, while errors from `resp.Body.Close()` are not common, it's a good practice to at least log them for debugging purposes. Ignoring them is often acceptable for simple applications, but more robust applications should consider logging or other forms of monitoring.