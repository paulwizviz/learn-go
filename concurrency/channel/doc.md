# Channels

A channel is a mechanism for goroutines and the main routine to communicate with each other. There are two types of channels:

* [Unbuffered channel](#unbuffered-channels)
* [Buffered channel](#buffered-channels)

## Unbuffered channels

This type of channel requires a receiving channel to be available when a message is emitted from a sending channel. If there is no receiver it will cause a deadlock.

The following examples are based on unbuffered channel.

* [Example 1a](./ex1a/main.go) we create an unbuffered channel without a receiving channel, causing Goroutine deadlock.
* [Example 1b](./ex1b/main.go) demonstrates blocking operations with channel
* [Example 1c](./ex1c/main.go) demonstrates a scenario with safeguard to check if channel is closed.
* [Example 1d](./ex1d/main.go) demonstrates a scenario with no safeguard.
* [Example 1e](./ex1e/main.go) demonstrates closing of channel prematurely.
* [Example 1f](./ex1f/main.go) using alternative ways to signal closure of channel.
* [Example 1g](./ex1g/main.go) using ranging to avoid goroutine deadlock

## Buffered channels

A buffer channel may receive message without any receiver. If you overflow the buffer it will cause a deadlock. If the receiver tries to access more message than is available in the buffer it will cause a deadlock.

The following examples are based on buffered channel.

* [Example 2a](./ex2a/main.go) demonstrates the use of buffered channel before a receiver is available and no deadlock (see also [Example 1a](./channel/ex1a/main.go) for opposite)
* [Example 2b](./ex2b/main.go) when you try to pull more messages than is available in the buffer
* [Example 2c](./ex2c/main.go) this cause a deadlock if you push more messages to the buffer's capacity.
* [Example 2d](./ex2d/main.go) this use select to prevent more message when buffer is full

