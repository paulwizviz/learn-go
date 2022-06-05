// Package http demonstrates the use of context in a http scenario
package http

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println("server: hello handler started")
	defer log.Println("server: hello handler ended")

	select {
	case <-time.After(5 * time.Second):
		// Wait until five seconds and then send response
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// If client end is cancelled before 5 seconds this will
		// be called
		err := ctx.Err()
		log.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
