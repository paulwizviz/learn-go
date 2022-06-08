package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func processConn(c net.Conn) {
	buf := make([]byte, 1024)
	nr, err := c.Read(buf)
	if err != nil {
		log.Fatalf("Unable to read. Reason: %v", err)
	}
	fmt.Printf("Received: %s\n", string(buf[:nr]))
	data := []byte{'h', 'e'}
	c.Write(data)
	c.Close()

}

func main() {
	pwd, _ := os.Getwd()
	socketFile := fmt.Sprintf("%s/temp/echo.sock", pwd)

	fmt.Println(socketFile)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(socketFile)
		os.Exit(1)
	}()

	l, err := net.Listen("unix", socketFile)
	if err != nil {
		log.Fatalf("Unable to start server. Reason: %s", err.Error())
	}
	defer os.Remove(socketFile)

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
		processConn(fd)
	}
}
