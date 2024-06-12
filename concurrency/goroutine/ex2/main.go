package main

import (
	"fmt"
	"time"
)

func main() {
	// This Go routine executes and hands off to second go routine.
	// It runs for ever until main routine ends.
	go func() {
		for {
			fmt.Println("Go routine 1")
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// This Go routine executes and hands off to main routine.
	// It runs forevet until main routine ends.
	go func() {
		for {
			fmt.Println("Go routine 2")
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Main routine sleeps for 1 second then ends.
	time.Sleep(1 * time.Second)
}
