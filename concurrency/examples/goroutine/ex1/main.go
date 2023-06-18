package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("Hello")
	}()

	fmt.Println("World")
}
