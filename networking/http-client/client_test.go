package client

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

func Example_get() {
	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			fmt.Println("Method is GET")
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("hello"))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	u := fmt.Sprintf("%s/%s", server.URL, "hello")

	req, _ := http.NewRequest(http.MethodGet, u, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	// Output:
	// Method is GET
	// hello
}

func Example_post() {
	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			fmt.Println("Method is POST")
		}
		body, _ := io.ReadAll(req.Body)
		fmt.Printf("Request Body: %s", body)
		rw.WriteHeader(http.StatusOK)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	u := fmt.Sprintf("%s/%s", server.URL, "hello")
	reader := strings.NewReader("{\"greetings\":\"hello\"}")
	req, _ := http.NewRequest(http.MethodPost, u, reader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Output:
	// Method is POST
	// Request Body: {"greetings":"hello"}
}
