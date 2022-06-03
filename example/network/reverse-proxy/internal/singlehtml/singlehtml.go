package singlehtml

import (
	"fmt"
	"log"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func Run() {
	port := os.Getenv("SINGLE_HTML_PORT")
	router := mux.NewRouter()

	htmlBox := rice.MustFindBox("../../web/singlehtml").HTTPBox()
	webContent := http.FileServer(htmlBox)

	router.PathPrefix("/").Handler(http.StripPrefix("/", webContent))

	log.Printf(fmt.Sprintf("Starting single html on Port: %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), router))
}
