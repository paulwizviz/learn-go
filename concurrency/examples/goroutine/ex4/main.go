package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var mutext *sync.Mutex = &sync.Mutex{}

	var value = 0
	fmt.Printf("Initial shared value: %v\n -- \n", value)

	go func() {
		for i := 0; i < 2; i++ {
			time.Sleep(10 * time.Millisecond)
			func() {
				mutext.Lock()
				fmt.Printf("1 Before increment: %v\n", value)
				value = value + 1
				fmt.Printf("1 After increment: %v\n --- \n", value)
				mutext.Unlock()
			}()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 2; i++ {
			time.Sleep(10 * time.Millisecond)
			func() {
				mutext.Lock()
				fmt.Printf("2 Before decrement: %v\n", value)
				value = value - 1
				fmt.Printf("2 After decrement: %v\n --- \n", value)
				mutext.Unlock()
			}()
		}
		wg.Done()
	}()

	wg.Wait()
}
