package funcs_test

import "fmt"

func fnByValue(arg int) {
	arg = 1
	fmt.Printf("Value of argument %v\n", arg)
}

func fnByRef(arg *int) {
	*arg = 1
	fmt.Printf("Value of argument %v\n", *arg)
}

func Example_mutatevalfunc() {

	value := 2
	fmt.Printf("Value before mutation %v\n", value)
	fnByValue(value)
	fmt.Printf("Value after mutation %v\n", value)

	fmt.Println("---")

	fmt.Printf("Value before mutation %v\n", value)
	fnByRef(&value)
	fmt.Printf("Value after mutation %v\n", value)

	fmt.Println("---")

	value1 := 5
	var intptr *int = &value1
	fmt.Printf("Value before mutation %v\n", *intptr)
	fnByRef(intptr)
	fmt.Printf("Value after mutation %v\n", *intptr)

	// Output:
	// Value before mutation 2
	// Value of argument 1
	// Value after mutation 2
	// ---
	// Value before mutation 2
	// Value of argument 1
	// Value after mutation 1
	// ---
	// Value before mutation 5
	// Value of argument 1
	// Value after mutation 1
}

func swapArray(arg [2]int) {
	arg[0], arg[1] = arg[1], arg[0]
}

func Example_mutatearrayfunc() {
	array := [2]int{1, 2}
	fmt.Printf("Array before mutation %v\n", array)
	swapArray(array)
	fmt.Printf("Array after mutation %v\n", array)

	// Output:
	// Array before mutation [1 2]
	// Array after mutation [1 2]
}

func swapSlice(arg []int) {
	arg[0], arg[1] = arg[1], arg[0]
}

func Example_mustateslicefunc() {
	slice := []int{1, 2}
	fmt.Printf("Slice before mutation %v\n", slice)
	swapSlice(slice)
	fmt.Printf("Slice after mutation %v\n", slice)

	// Output:
	// Slice before mutation [1 2]
	// Slice after mutation [2 1]
}
