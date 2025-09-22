package main

import (
	"fmt"
	"log"
	"net"
)

func udpServer(serverAddr string) (net.Addr, error) {
	server, err := net.ListenPacket("udp", serverAddr)
	if err != nil {
		return nil, err
	}

	go func(c net.PacketConn, sa string) {
		b := make([]byte, 1024)
		for {
			n, clientAddr, err := c.ReadFrom(b) // Message from client
			if err != nil {
				fmt.Println(err)
				break
			}

			s := string(b[:n]) + " World"
			b := []byte(s)
			_, err = c.WriteTo(b, clientAddr) // Message to client
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}(server, serverAddr)
	return server.LocalAddr(), nil
}

func main() {
	sAddr, err := udpServer("127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}

	client, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	msg := []byte("hello")
	fmt.Printf("Message from; %v message: %v\n", client.LocalAddr(), string(msg))
	_, err = client.WriteTo(msg, sAddr) // Write to server
	if err != nil {
		log.Fatal(err)
	}
	b := make([]byte, 1024)
	n, serverAddr, err := client.ReadFrom(b) // Receive from server
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Receive from: %v message: %v", serverAddr, string(b[:n]))
}
