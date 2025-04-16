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
	defer func() {
		// Read the response body to ensure it is fully consumed.
		// This is important to prevent resource leaks.
		// In a real-world scenario, you would typically read the body
		// and process it, but here we just want to demonstrate
		// that we can read it without any issues.
		io.Copy(io.Discard, resp.Body)
		// Close the response body to prevent resource leaks
		// and to ensure the server can reuse the connection.
		// This is important in a real-world scenario.
		// In this example, we are using a test server,
		// so it is not strictly necessary, but it's a good practice.
		err := resp.Body.Close()
		if err != nil {
			// This is not necessary in a test server,
			// but in a real-world scenario, you should handle errors.
			// For example, you might want to log the error or return it.
			fmt.Println("Error closing response body:", err)
		}
	}()

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
