package deferops

import "fmt"

func Example_defer() {
	defer fmt.Println("World")
	fmt.Println("Hello")

	// Output:
	// Hello
	// World
}

func Example_nestedDefer() {
	defer fmt.Println("World")
	defer fmt.Println("Mundo")

	fmt.Println("Hello")
	fmt.Println("Ola")

	// Output:
	// Hello
	// Ola
	// Mundo
	// World

}

func Example_midDefer() {
	defer fmt.Println("World")

	if true {
		fmt.Println("No Hello")
		return
	}
	fmt.Println("Hello")

	// Output:
	// No Hello
	// World

}

func errorCalled() error {
	defer fmt.Println("Hello")
	if true {
		return fmt.Errorf("error happens")
	}
	return nil
}

func Example_errFnCalled() {
	defer fmt.Println("World") // Defer to end of example called
	err := errorCalled()
	fmt.Println(err)

	// Output:
	// Hello
	// error happens
	// World

}

func noErrorCalled() error {
	defer fmt.Println("Hello")
	if false {
		return fmt.Errorf("error happens")
	}
	return nil
}

func Example_noErrFnCalled() {
	defer fmt.Println("World") // Defer to end of example called
	err := noErrorCalled()
	fmt.Println(err)

	// Output:
	// Hello
	// <nil>
	// World
}
