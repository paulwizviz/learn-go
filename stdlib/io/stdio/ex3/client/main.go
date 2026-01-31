package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// 1. Prepare the command to run the server binary
	cmd := exec.Command("go", "run", "../server/main.go")

	// 2. Attach pipes to the server's Stdin and Stdout
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	// 3. Start the process
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// 4. Send a message to the server
	go func() {
		// This will cause the server side scanner.Scan to return false
		defer stdin.Close()
		fmt.Fprintln(stdin, "hello from client")
		fmt.Fprintln(stdin, "stdio is efficient")
	}()

	// 5. Read responses from the server
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Printf("Received: %s\n", scanner.Text())
	}

	// 6. Wait for the server to exit, the main routine is blocked
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
