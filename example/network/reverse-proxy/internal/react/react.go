package react

import (
	"fmt"
	"log"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

// Process to start a web server
func Run() {
	port := os.Getenv("REACT_PORT")
	router := mux.NewRouter()

	box := rice.MustFindBox("../../web/public")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(box.HTTPBox())))

	log.Printf(fmt.Sprintf("Starting web server on Port: %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), router))
}
