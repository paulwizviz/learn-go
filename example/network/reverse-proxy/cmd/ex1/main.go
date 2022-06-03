package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/paulwizviz/learn-go/example/network/reverse-proxy/internal/proxy"
	"github.com/paulwizviz/learn-go/example/network/reverse-proxy/internal/react"
	"github.com/paulwizviz/learn-go/example/network/reverse-proxy/internal/singlehtml"

	"github.com/paulwizviz/learn-go/example/network/reverse-proxy/internal/api"
)

func setup() {
	// Ports hardcoded for illustrative purpose
	// In a more production like system, you should be
	// reading the ports from environment variables
	os.Setenv("PROXY_PORT", "80")
	os.Setenv("REACT_PORT", "3000")
	os.Setenv("SINGLE_HTML_PORT", "3010")
	os.Setenv("API_PORT", "9000")
}

func main() {
	setup()
	args := os.Args

	if len(args) != 2 {
		log.Fatalf("Usage: %v [simple|react]", filepath.Base(args[0]))
	}

	os.Setenv("WEB_TYPE", args[1])
	if args[1] == "simple" {
		go func() {
			singlehtml.Run()
		}()
	} else if args[1] == "react" {
		go func() {
			react.Run()
		}()
	} else {
		log.Fatalf("Usage: %v [simple|react]", filepath.Base(args[0]))
	}
	go func() {
		api.Run()
	}()
	proxy.Run()
}
