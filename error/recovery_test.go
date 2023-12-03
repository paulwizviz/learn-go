package examples

import "fmt"

func div(x, y int) {
	fmt.Println(x / y)
}

func Example_recovery() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	fmt.Println("Before divide")
	div(1, 0)
	fmt.Println("After divide") // This will never be called

	// Output:
	// Before divide
	// runtime error: integer divide by zero
}

func math() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	fmt.Println("Before divide")
	div(1, 0)
	fmt.Println("After divide") // This will never be called
}

func Example_math() {
	fmt.Println("Before math")
	math()
	fmt.Println("After math")

	// Output:
	// Before math
	// Before divide
	// runtime error: integer divide by zero
	// After math
}

func mathNoDefer() {
	fmt.Println("Before divide")
	div(1, 0)
	fmt.Println("After divide") // This will never be called
}

func Example_mathNoDefer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	fmt.Println("Before math")
	mathNoDefer()
	fmt.Println("After math") // This will never be called

	// Output:
	// Before math
	// Before divide
	// runtime error: integer divide by zero

}
