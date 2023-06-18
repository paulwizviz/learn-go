package os_test

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Example_unixEchoCommand() {
	cmd := exec.Command("echo", "hello")
	var outputBuff bytes.Buffer
	cmd.Stdout = &outputBuff
	if err := cmd.Run(); err != nil {
		fmt.Printf("Unable to run command. Reason: %v", err)
	}

	fmt.Println(outputBuff.String())
	// Output:
	// hello
}

func Example_unixTranslate() {
	cmd := exec.Command("tr", "a-z", "A-Z") // translate any characters from lower to upper case
	// Command will stop and take input characters
	cmd.Stdin = strings.NewReader("version")
	var outBuff bytes.Buffer
	cmd.Stdout = &outBuff
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Unable to run command. Reason: %v", err)
	}
	fmt.Println(outBuff.String())
	// Output:
	// VERSION
}
