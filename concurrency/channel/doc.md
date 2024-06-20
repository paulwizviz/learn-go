# Channels

A channel is a mechanism for goroutines to communicate with each other. Think of channels like a pipe where messages are sent.

There are two types of channels:

* [Unbuffered channel](#unbuffered-channels)
* [Buffered channel](#buffered-channels)

## Unbuffered channels

This type of channel requires a receiving channel to be available when a message is emitted from a sending channel. If there is no receiver it will cause a deadlock.

### Example 1a - Sender and receiver

These examples demonstrate the relationship between sender and receiver channels.

* [Example 1a1](./ex1a1/main.go) shows a scenario where a sender has sent a message without a receiver.
* [Example 1a2](./ex1a2/main.go) this fix the problem of `Example 1a1`.

### Example 1b - Blocking operations

These examples demonstrate channel blocking operations.

* [Example 1b1](./ex1b1/main.go) demonstrates a blocked sending channel.
* [Example 1b2](./ex1b2/main.go) demonstrates a blocked receiver channel.

### Example 1c - Closing sender

In many real world scenarios, multiple messages sent to a sender. However, the receiver does not know when the sender has stopped sending messages. 

* [Example 1c1](./ex1c1/main.go) this scenario shows receiver attempting to receive more messages than is available, thus causing a deadlock.
* [Example 1c2](./ex1c2/main.go) demonstrates closing a sender channel to signal to reciever to not expect any more messages.
* [Example 1c3](./ex1c3/main.go) demonstrates closing of channel prematurely.
* [Example 1c4](./ex1c4/main.go) using alternative ways to signal closure of channel.
* [Example 1c5](./ex1c5/main.go) using ranging to avoid goroutine deadlock

## Buffered channels

A buffer channel may receive message without any receiver. If you overflow the buffer it will cause a deadlock. If the receiver tries to access more message than is available in the buffer it will cause a deadlock.

### Example 2a - Sender and receiver

These examples demonstrates the relationship between sender and receiver.

* [Example 2a1](./ex2a1/main.go) demonstrates the use of buffered channel before a receiver is available and no deadlock (see also [Example 1a](./channel/ex1a/main.go) for opposite).
* [Example 2a2](./ex2a2/main.go) demonstrate a deadlock situation when you try to push more to than the buffer capacity .

### Example 2b - Blocking operations

These examples show operations where channels are blocked.

* [Example 2b1](./ex2b1/main.go) this demonstrates channel is blocked when more messages is sent than capacity.
* [Example 2b2](./ex2d/main.go) this use select to prevent more message pushed to sender when buffer is full

### Example 2c - Synchronising between goroutine

These examples show the use of signals to synchronise channels across goroutine.

* [Example 2c1](./ex2c1/main.go) demonstrates lack of safeguards to prevent deadlock.
* [Example 2c2](./ex2c2/main.go) demionstrates the use of range to safeguard channels.
* [Example 2c3](./ex2c3/main.go) demonstrates safeguard to ensure loop breaks when signal close is trigger. 

