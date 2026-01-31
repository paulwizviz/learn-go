package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	// Set up an external command
	cmd := exec.Command("ping", "google.com")

	// Obtain a standard out pipeline
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// Exeute command
	cmd.Start()

	// Create a buffered reader
	reader := bufio.NewReader(stdout)
	lineCount := 0 // initialise line count

	for {
		line, _, _ := reader.ReadLine()
		if lineCount > 3 {
			os.Exit(0)
		}
		lineCount += 1
		fmt.Println(string(line))
	}
}
