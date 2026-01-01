package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // Count two wait groups

	go func() {
		fmt.Println("Goroutine 1")
		wg.Done() // Wait done - decrease by one
	}()

	go func() {
		fmt.Println("Goroutine 2")
		wg.Done() // Wait done - decrease by one
	}()
	fmt.Println("Number of Goroutine before wait: ", runtime.NumGoroutine()) // This will report 3 goroutines (main is also a goroutine)
	wg.Wait()                                                                // Cause the main routine to pause until all routine is done
	fmt.Println("Number of Goroutine after wait: ", runtime.NumGoroutine())  // This will report 1 goroutine (main routine)
}
