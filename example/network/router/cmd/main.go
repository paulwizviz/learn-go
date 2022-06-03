package main

import (
	"fmt"
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func webServer() {
	port := "3030"
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("../web").HTTPBox()))
	log.Printf("Starting react %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), router))
}

func main() {
	webServer()
}
