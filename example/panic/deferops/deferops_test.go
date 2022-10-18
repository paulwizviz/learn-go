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

func errorCalled() error {
	defer fmt.Println("world")
	if true {
		fmt.Println("Error occured")
		return fmt.Errorf("error happens")
	}
	return nil
}

func noErrorCalled() error {
	defer fmt.Println("world")
	if false {
		fmt.Println("Error occured")
		return fmt.Errorf("error happens")
	}
	return nil
}

func Example_deferCalledAfterEtn() {
	err := errorCalled()
	fmt.Println(err)

	fmt.Println("---")

	err = noErrorCalled()
	fmt.Println(err)

	// Output:
	// Error occured
	// world
	// error happens
	// ---
	// world
	// <nil>

}
