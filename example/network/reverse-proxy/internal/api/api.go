package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// A simple rest endpoint handler that simply return a phrase hello
func homeHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("123456")
	res.WriteHeader((http.StatusOK))
	fmt.Fprint(res, "{hello:hello}")
}

// Process to start the REST server
func Run() {
	port := os.Getenv("API_PORT")
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	log.Printf("Starting API on port: %s", port)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), router)
}
