package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"
	"time"
)

// client operation dials a connection
func client(socket string, msg string) error {
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func processConn(c net.Conn) {
	buf := make([]byte, 1024)
	nr, err := c.Read(buf)
	if err != nil {
		log.Fatalf("Unable to read. Reason: %v", err)
	}
	fmt.Printf("Received: %s\n", string(buf[:nr]))
}

func server(socket string) error {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(socket)
		os.Exit(0)
	}()

	// socket listening
	l, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatalf("Unable to start server. Reason: %s", err.Error())
	}
	defer os.Remove(socket)

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
		processConn(fd)
	}
}

func main() {

	msg := flag.String("message", "", "Message")
	flag.Parse()

	if *msg == "" {
		log.Fatal("message flag missing")
	}

	pwd, _ := os.Getwd()
	socket := path.Join(pwd, "temp", "socket")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		err := server(socket)
		if err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	time.Sleep(1 * time.Second)

	err := client(socket, *msg)
	if err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
