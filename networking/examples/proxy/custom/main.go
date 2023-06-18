package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

type proxyHandler struct{}

func (p *proxyHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var forwardURL *url.URL
	var err error

	if req.URL.Path == "/" {
		forwardPort := 3031
		forwardURL, err = url.ParseRequestURI(fmt.Sprintf("http://localhost:%v/", forwardPort))
		if err != nil {
			log.Fatal("Unable to start a server")
		}
		// req.RemoteAddr is from the originating address
		log.Printf("Remote address: %v", req.RemoteAddr)
		req.Host = forwardURL.Host
		req.URL.Host = forwardURL.Host
		req.URL.Scheme = forwardURL.Scheme
		req.RequestURI = ""

		remoteHost, remotePort, _ := net.SplitHostPort(req.RemoteAddr)
		remoteURL := fmt.Sprintf("%v:%v", remoteHost, remotePort)
		req.Header.Set("X-Forwarded-For", remoteURL)

		response, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(rw, err)
		}

		rw.WriteHeader(response.StatusCode)
		io.Copy(rw, response.Body)
	}

}

func main() {

	port := 3030
	http.Handle("/", &proxyHandler{})

	server := &http.Server{
		Addr:        fmt.Sprintf("0.0.0.0:%v", port),
		ReadTimeout: time.Second * 1,
	}
	log.Printf("Starting proxy on %v", port)
	log.Fatal(server.ListenAndServe())
}
