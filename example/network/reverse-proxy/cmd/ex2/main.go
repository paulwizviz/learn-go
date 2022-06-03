package main

import (
	"os"

	"github.com/paulwizviz/learn-go/example/network/reverse-proxy/internal/customproxy"
	"github.com/paulwizviz/learn-go/example/network/reverse-proxy/internal/react"
)

func setup() {
	os.Setenv("PROXY_PORT", "80")
	os.Setenv("REACT_PORT", "8001")
}

func main() {
	setup()
	go func() {
		react.Run()
	}()
	customproxy.Run()
}
