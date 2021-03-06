package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "3030"
	// wd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatalf("Unable to obtain working directory, %v", err)
	// }
	dir := http.Dir("./")
	http.Handle("/", http.FileServer(dir))
	log.Printf("Starting webserver on %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
