package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMux(t *testing.T) {

	testcases := []struct {
		method     string
		urlPattern string
		urlTarget  string
		handler    http.HandlerFunc
		failed     bool
	}{
		{
			method:     http.MethodGet,
			urlPattern: "/test",
			urlTarget:  "/test",
			handler:    func(rw http.ResponseWriter, r *http.Request) {},
			failed:     false,
		},
		{
			method:     http.MethodGet,
			urlPattern: "/test/{id}",
			urlTarget:  "/test/123",
			handler: func(rw http.ResponseWriter, r *http.Request) {
				id := r.PathValue("id")
				t.Logf("Case: 1 id: %s", id)
			},
			failed: false,
		},
		{
			method:     http.MethodGet,
			urlPattern: "/test/{id}",
			urlTarget:  "/test/123/qw",
			handler: func(rw http.ResponseWriter, r *http.Request) {
				// The handler is ignored so the following part is
				// not executed
				id := r.PathValue("id")
				t.Logf("Case 2 id: %s", id)
			},
			failed: true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("Case-%d", i), func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.urlTarget, nil)
			rec := httptest.NewRecorder()
			mux := http.NewServeMux()
			mux.Handle(tc.urlPattern, tc.handler)
			mux.ServeHTTP(rec, req)
			if !tc.failed && rec.Code != http.StatusOK {
				t.Errorf("Case: %d Want: %d Got: %d", i, 200, rec.Code)
			}
			if tc.failed && rec.Code != http.StatusNotFound {
				t.Errorf("Case: %d Want: %d Got: %d", i, 404, rec.Code)
			}
		})
	}

}
