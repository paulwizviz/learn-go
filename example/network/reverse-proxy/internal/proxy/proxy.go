package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gorilla/mux"
)

// Handler responsible for routing path / to web server
func handleSingleHTMLRequest(res http.ResponseWriter, req *http.Request) {

	// Create the forwarding request URL
	forwardURL, err := url.ParseRequestURI(fmt.Sprintf("http://localhost:%v/", os.Getenv("SINGLE_HTML_PORT")))
	if err != nil {
		log.Fatalf("Failed to form forwarding URL: %v", err)
	}

	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.URL = forwardURL

	proxy := httputil.NewSingleHostReverseProxy(forwardURL)
	proxy.ServeHTTP(res, req)
}

func handleReactRequest(res http.ResponseWriter, req *http.Request) {

	// Create the forwarding request URL
	forwardURL, err := url.ParseRequestURI(fmt.Sprintf("http://localhost:%v/", os.Getenv("REACT_PORT")))
	if err != nil {
		log.Fatalf("Failed to form forwarding URL: %v", err)
	}

	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.URL = forwardURL

	proxy := httputil.NewSingleHostReverseProxy(forwardURL)
	proxy.ServeHTTP(res, req)
}

func handleAPIRequest(res http.ResponseWriter, req *http.Request) {

	forwardURL, err := url.ParseRequestURI(fmt.Sprintf("http://localhost:%v/", os.Getenv("API_PORT")))
	if err != nil {
		log.Fatalf("Failed to form forwarding URL: %v", err)
	}

	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.URL = forwardURL

	proxy := httputil.NewSingleHostReverseProxy(forwardURL)
	proxy.ServeHTTP(res, req)
}

// Process to start the proxy
func Run() {
	port := os.Getenv("PROXY_PORT")
	router := mux.NewRouter()

	webType := os.Getenv("WEB_TYPE")
	if webType == "simple" {
		router.HandleFunc("/", handleSingleHTMLRequest)
	} else if webType == "react" {
		router.HandleFunc("/", handleReactRequest)
	}
	router.HandleFunc("/api", handleAPIRequest)
	log.Printf("Proxy starting on %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), router))
}
