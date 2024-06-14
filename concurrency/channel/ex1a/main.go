package main

func main() {

	// Create an unbuffered channel
	c := make(chan string)
	// Send message to channel before receiver is available
	c <- "Hello World" // Causing Go routine dead lock.
	<-c
}
