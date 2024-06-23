# Channels

A channel is a mechanism for goroutines to communicate with each other. A channel is like a pipe; at each end are a sender and a receiver for sending and receiving messages respectively.

There are two types of channels:

* [Unbuffered channel](#unbuffered-channels)
* [Buffered channel](#buffered-channels)

## Unbuffered channels

This type of channel requires for both sender and receiver to be ready.

### Example 1a - Sender and receiver

These examples demonstrate the relationship between sender and receiver channels.

* [Example 1a1](./ex1a1/main.go) shows a scenario where a sender has sent a message without a receiver.
* [Example 1a2](./ex1a2/main.go) this fix the problem of `Example 1a1`.

### Example 1b - Blocking operations

These examples demonstrate channel blocking operations.

* [Example 1b1](./ex1b1/main.go) demonstrates a blocked sending channel.
* [Example 1b2](./ex1b2/main.go) demonstrates a blocked receiver channel.

### Example 1c - Synchronising channels

In many real world scenarios, multiple messages sent to a sender. However, the receiver does not know when the sender has stopped sending messages. 

* [Example 1c1](./ex1c1/main.go) this scenario shows receiver attempting to receive more messages than is available, thus causing a deadlock.
* [Example 1c2](./ex1c2/main.go) demonstrates closing a sender channel signalling to receiver not to expect any more messages.
* [Example 1c3](./ex1c3/main.go) demonstrates closing of channel prematurely, thus causing sender to panic.
* [Example 1c4](./ex1c4/main.go) using alternative ways to signal closure of channel.
* [Example 1c5](./ex1c5/main.go) using ranging to avoid goroutine deadlock

## Buffered channels

A buffered channel has a capacity to hold messages before it is drained. This will allow sender to send messages before receiver is ready.

### Example 2a - Sender and receiver

These examples demonstrates the relationship between sender and receiver.

* [Example 2a1](./ex2a1/main.go) demonstrates a buffered channel before a receiver is available.
* [Example 2a2](./ex2a2/main.go) demonstrates a deadlock situation when you try to push more messages than the buffer capacity.

### Example 2b - Blocking operations

This [Example 2b](./ex2b/main.go) shows operation when the buffered channel if filled to capacity before it is drained. The channel is blocked and not able to receive more messages. 

### Example 2c - Synchronising between goroutine

These examples show the use of signals to synchronise channels across goroutine.

* [Example 2c1](./ex2c1/main.go) demonstrates receiver draining more messages than is available without safeguards to prevent deadlock.
* [Example 2c2](./ex2c2/main.go) demionstrates the use of range to safeguard channels.
* [Example 2c3](./ex2c3/main.go) demonstrates safeguard to ensure loop breaks when signal close is trigger. 
* [Example 2c4](./ex2c4/main.go) demonstrates the use of buffered function for asynchronous function.

