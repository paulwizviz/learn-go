package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Read line-by-line from Stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() { // This will return false when client stdin.Close() is called
		input := scanner.Text()

		// 2. Process logic (e.g., convert to uppercase)
		result := strings.ToUpper(input)

		// 3. Write response to Stdout with a newline delimiter
		fmt.Fprintln(os.Stdout, result)

		// Optional: Log to Stderr (visible to the parent process)
		fmt.Fprintln(os.Stderr, "Server processed message: "+input)
	}
}
