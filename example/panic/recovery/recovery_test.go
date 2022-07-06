package main

import "fmt"

func Example_recovery() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered")
		}
	}()
	fmt.Println("Before panic")
	panic("Panic called")
	fmt.Println("After panic") // This will never be called

	// Output:
	// Before panic
	// Recovered
}

func panickingFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panicking func Recovered")
		}
	}()
	fmt.Println("Panicking func called")
	panic("Func panic")
	fmt.Println("Panicking func end")
}

func toplevelfunc() {
	fmt.Println("Top level func called")
	panickingFunc() // Calling another function that is potentially panicking
	fmt.Println("Top level func end")
}

// In this example, we want to stop a potentially
// panicking func from panicking
func Example_preventFailing() {
	toplevelfunc()
	// Output:
	// 	Top level func called
	// Panicking func called
	// Panicking func Recovered
	// Top level func end
}
