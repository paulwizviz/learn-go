package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Goroutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Goroutine 2")
		wg.Done()
	}()
	fmt.Println("Number of Goroutine before wait: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Number of Goroutine after wait: ", runtime.NumGoroutine())
}
