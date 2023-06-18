package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello, 世界")
	}()

	time.Sleep(100 * time.Microsecond)
}
