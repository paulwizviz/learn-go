package main

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
)

func ifConfig() string {
	cmd := exec.Command("ifconfig", "-a")
	var outputBuff bytes.Buffer
	cmd.Stdout = &outputBuff
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
	return outputBuff.String()

}

func networkInterface() string {
	out := ""
	nis, _ := net.Interfaces()
	for _, ni := range nis {
		addrs, err := ni.Addrs()
		if err != nil {
			fmt.Printf("No addresses: %v", err)
			break
		}
		for _, addr := range addrs {
			out = out + fmt.Sprintf("Name: %v Address: %v\n", ni.Name, addr)
		}
	}
	return out
}

func assignPort() string {
	// Port :0 causes the listener to assign a port number
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("Port assigned: %v", listener.Addr().(*net.TCPAddr).Port)
}

func main() {
	fmt.Println("Result from ifconfig")
	fmt.Println(ifConfig())
	fmt.Println("======")
	fmt.Println("Result from net.Interface")
	fmt.Println(networkInterface())
	fmt.Println("======")
	fmt.Printf("Assigned poer: %v", assignPort())
}
