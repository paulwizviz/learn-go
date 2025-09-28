package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("Starting tasks...")

	// With Go 1.25, wg.Go handles Add(1) and Done() automatically.
	for i := 0; i < 3; i++ {
		id := i // Capture the loop variable for the closure.
		wg.Go(func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Task %d finished\n", id)
		})
	}

	wg.Wait() // Wait for all goroutines started with wg.Go to finish.

	fmt.Println("All tasks completed.")
}
