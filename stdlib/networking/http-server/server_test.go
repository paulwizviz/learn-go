package server

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// This demonstrates the use of http package based
// on standard library
func TestStdMux(t *testing.T) {

	testcases := []struct {
		method         string
		urlPattern     string
		urlTarget      string
		handler        http.HandlerFunc
		wantBody       []byte
		wantStatusCode int
	}{
		{
			method:         http.MethodGet,
			urlPattern:     "/test",
			urlTarget:      "/test",
			handler:        func(rw http.ResponseWriter, r *http.Request) {},
			wantBody:       nil,
			wantStatusCode: http.StatusOK,
		},
		{
			method:     http.MethodGet,
			urlPattern: "/test/{id}",
			urlTarget:  "/test/123",
			handler: func(rw http.ResponseWriter, r *http.Request) {
				id := r.PathValue("id")
				rw.Write([]byte(id))
			},
			wantBody:       []byte("123"),
			wantStatusCode: http.StatusOK,
		},
		{
			method:     http.MethodGet,
			urlPattern: "/test/{id}",
			urlTarget:  "/test/123/qw",
			handler: func(rw http.ResponseWriter, r *http.Request) {
				// The handler is not call as the request url is never match
				id := r.PathValue("id")
				rw.Write([]byte(id))
			},
			wantBody:       []byte("404 page not found\n"),
			wantStatusCode: http.StatusNotFound,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("Case-%d", i), func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.urlTarget, nil)
			rec := httptest.NewRecorder()
			mux := http.NewServeMux()
			mux.Handle(tc.urlPattern, tc.handler)
			mux.ServeHTTP(rec, req)
			gotBody := rec.Body.Bytes()
			t.Log(tc.wantBody, gotBody)
			if rec.Code != tc.wantStatusCode || !bytes.Equal(tc.wantBody, gotBody) {
				t.Errorf("Code want: %d, got: %d", tc.wantStatusCode, rec.Code)
				t.Errorf("Message want: %s, got: %s", tc.wantBody, gotBody)
			}
		})
	}
}

type HttpHandler struct {
	mu      sync.Mutex
	message string
}

func (h *HttpHandler) Handler1(rw http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	rw.Write([]byte(h.message))
}
func (h *HttpHandler) Handler2(rw http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	id := r.PathValue("id")
	msg := fmt.Sprintf("%s-%s", h.message, id)
	rw.Write([]byte(msg))
}

func TestHttpHandler(t *testing.T) {
	testcases := []struct {
		method    string
		urlTarget string
		want      []byte
	}{
		{
			method:    http.MethodGet,
			urlTarget: "/hello",
			want:      []byte("calling database"),
		},
		{
			method:    http.MethodGet,
			urlTarget: "/hello/123",
			want:      []byte("calling database-123"),
		},
		{
			method:    http.MethodGet,
			urlTarget: "/hello/123/world", // Target never match
			want:      []byte("404 page not found\n"),
		},
	}

	for i, tc := range testcases {
		hdlr := &HttpHandler{
			message: "calling database",
		}
		t.Run(fmt.Sprintf("case: %d", i), func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.urlTarget, nil)
			rec := httptest.NewRecorder()
			mux := http.NewServeMux()
			mux.HandleFunc("/hello", hdlr.Handler1)
			mux.HandleFunc("/hello/{id}", hdlr.Handler2)
			mux.ServeHTTP(rec, req)
			got, _ := io.ReadAll(rec.Body)
			if !bytes.Equal(got, tc.want) {
				t.Fatalf("Want: %s Got: %s", tc.want, got)
			}
		})
	}
}
