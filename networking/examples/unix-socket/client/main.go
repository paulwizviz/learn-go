package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	socketFile := fmt.Sprintf("%s/temp/echo.sock", pwd)
	conn, err := net.Dial("unix", socketFile)
	if err != nil {
		log.Fatal(err)
	}
	msg := []byte("hello")
	_, err = conn.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
}
