package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Random port assigned
	l, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}
	go func(l net.Listener) {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println(err)
				break
			}
			b := make([]byte, 1024)
			for {
				n, err := conn.Read(b)
				if err == io.EOF {
					fmt.Println(err)
					break
				}
				fmt.Println(string(b[:n]))
			}
		}
	}(l)

	conn, err := net.Dial("tcp", l.Addr().String())
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte("Hello"))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")

	conn.Close()
	l.Close()
}
