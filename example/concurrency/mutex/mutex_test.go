package mutex_test

import (
	"fmt"
	"sync"
	"time"
)

func Example_shareVariableNoMutex() {

	var wg sync.WaitGroup
	wg.Add(2)

	var value = 0

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 100)
			func() {
				fmt.Printf("1 Before: %v\n", value)
				value = value + 1
				fmt.Printf("1 After: %v\n", value)
			}()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 100)
			func() {
				fmt.Printf("2 Before: %v\n", value)
				value = value - 1
				fmt.Printf("2 After: %v\n", value)
			}()
		}
		wg.Done()
	}()

	wg.Wait()
}

func Example_shareVariableWithMutex() {

	var wg sync.WaitGroup
	wg.Add(2)

	var mutext *sync.Mutex = &sync.Mutex{}

	var value = 0

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 100)
			func() {
				mutext.Lock()
				fmt.Printf("1 Before: %v\n", value)
				value = value + 1
				fmt.Printf("1 After: %v\n", value)
				mutext.Unlock()
			}()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 100)
			func() {
				mutext.Lock()
				fmt.Printf("2 Before: %v\n", value)
				value = value - 1
				fmt.Printf("2 After: %v\n", value)
				mutext.Unlock()
			}()
		}
		wg.Done()
	}()

	wg.Wait()
}
