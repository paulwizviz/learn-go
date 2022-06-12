package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func redirectRequest(res http.ResponseWriter, req *http.Request) {
	targetURL, _ := url.Parse("http://localhost:3031")

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	req.URL.Host = targetURL.Host
	req.URL.Scheme = targetURL.Scheme
	req.Header.Set("X-Forward-Host", req.Header.Get("Host"))
	req.Host = targetURL.Host

	proxy.ServeHTTP(res, req)

}

func main() {
	port := 3030
	http.HandleFunc("/", redirectRequest)

	log.Printf("Starting server at port: %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
