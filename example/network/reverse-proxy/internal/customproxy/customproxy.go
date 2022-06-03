package customproxy

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type proxyHandler struct{}

func (p *proxyHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	var forwardURL *url.URL
	var err error

	if req.URL.Path == "/" {
		forwardURL, err = url.ParseRequestURI(fmt.Sprintf("http://localhost:%v/", os.Getenv("REACT_PORT")))
		if err != nil {
			log.Fatal("Unable to start a server")
		}
		log.Println(req.RemoteAddr)
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

		for key, values := range response.Header {
			for _, v := range values {
				rw.Header().Set(key, v)
			}
		}

		done := make(chan bool)
		go func() {
			for {
				select {
				case <-time.Tick(1 * time.Millisecond):
					rw.(http.Flusher).Flush()
				case <-done:
					return
				}
			}
		}()

		trailerKeys := []string{}
		for key := range response.Trailer {
			trailerKeys = append(trailerKeys, key)
		}

		rw.Header().Set("Trailer", strings.Join(trailerKeys, ","))

		rw.WriteHeader(response.StatusCode)
		io.Copy(rw, response.Body)

		for key, values := range response.Trailer {
			for _, value := range values {
				rw.Header().Set(key, value)
			}
		}
		close(done)
	}

}

func Run() {
	port := os.Getenv("PROXY_PORT")

	http.Handle("/", &proxyHandler{})

	server := &http.Server{
		Addr:        fmt.Sprintf("0.0.0.0:%v", port),
		ReadTimeout: time.Second * 1,
	}
	log.Printf("Starting proxy on %v", port)
	log.Fatal(server.ListenAndServe())
}
