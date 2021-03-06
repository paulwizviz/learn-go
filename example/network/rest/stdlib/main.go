package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println("--- root ----")
		log.Printf("Request Host: %v", req.Host)
		log.Printf("Request header %v", req.Header)
		log.Printf("Request requestURI %v", req.RequestURI)
		log.Printf("Request URL (string): %v", req.URL.String())
		log.Printf("Request URL.Scheme %v", req.URL.Scheme)
		log.Printf("Request URL.Path: %v", req.URL.Path)
		log.Printf("Request URL.Hostname: %v", req.URL.Hostname())
		log.Printf("Request URL.RequestURI: %v", req.URL.RequestURI())
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("Hello"))
	})

	http.HandleFunc("/ola", func(res http.ResponseWriter, req *http.Request) {
		log.Println("--- /ola ----")
		log.Printf("Request Host: %v", req.Host)
		log.Printf("Request header %v", req.Header)
		log.Printf("Request requestURI %v", req.RequestURI)
		log.Printf("Request URL (string): %v", req.URL.String())
		log.Printf("Request URL.Scheme %v", req.URL.Scheme)
		log.Printf("Request URL.Path: %v", req.URL.Path)
		log.Printf("Request URL.Hostname: %v", req.URL.Hostname())
		log.Printf("Request URL.RequestURI: %v", req.URL.RequestURI())
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("ola"))
	})

	log.Fatal(http.ListenAndServe(":3030", nil))
}
