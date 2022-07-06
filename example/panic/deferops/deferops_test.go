package deferops

import "fmt"

func Example_defer() {
	defer fmt.Println("World")
	fmt.Println("Hello")

	// Output:
	// Hello
	// World
}

func Example_deferErrCalled() {
	defer fmt.Println("World")
	if true {
		fmt.Println("Error called")
		return
	}
	fmt.Println("Hello")

	// Output:
	// Error called
	// World
}

func Example_deferNoErr() {
	defer fmt.Println("World")
	if false {
		fmt.Println("Error called")
	}
	fmt.Println("Hello")

	// Output:
	// Hello
	// World
}
