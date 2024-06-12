package main

func main() {

	// Create an unbuffered channel
	c := make(chan string)
	// Send message to channel before receiver
	c <- "Hello World"
	// Receivers tries to get message
	<-c // Causing Go routine dead lock.
}
