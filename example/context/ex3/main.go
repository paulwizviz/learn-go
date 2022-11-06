package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fmt.Println("processing request")
		select {
		case <-time.After(2 * time.Second):
			w.Write([]byte("request processed"))
		case <-ctx.Done():
			fmt.Println("request cancelled")
		}
	}))
}
