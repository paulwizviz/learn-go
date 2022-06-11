package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Printf("Request Host: %v", req.Host)
		log.Printf("Request requestURI %v", req.RequestURI)
		log.Printf("Request URL (string): %v", req.URL.String())
		res.WriteHeader(http.StatusOK)
		reply := fmt.Sprintf("Hello from %v", req.Host)
		res.Write([]byte(reply))
	})
	port := "3031"
	log.Printf("Starting server at port: %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
