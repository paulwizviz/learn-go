package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context() // This prints to STDOUT to show that processing has started
		fmt.Println("processing request")
		// We use `select` to execute a piece of code depending on which
		// channel receives a message first
		select {
		case <-time.After(2 * time.Second):
			// If we receive a message after 2 seconds
			// that means the request has been processed
			// We then write this as the response
			w.Write([]byte("request processed"))
		case <-ctx.Done():
			// If the request (e.g. browser closes) is cancelled in less than 2 secs
			fmt.Println("request cancelled")
		}
	}))
}
