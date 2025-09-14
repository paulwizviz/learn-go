package types

import "fmt"

func fnByValue(arg int) {
	arg = 1 // mutate argument
	fmt.Printf("Value of argument %v\n", arg)
}

func Example_fncByValue() {
	value := 2
	fmt.Printf("Value before mutation %v\n", value)
	fnByValue(value)
	fmt.Printf("Value after mutation %v\n", value)

	// Output:
	// Value before mutation 2
	// Value of argument 1
	// Value after mutation 2
}

func fnByRef(arg *int) {
	*arg = 1
	fmt.Printf("Value of argument %v\n", *arg)
}

func Example_fnByRef() {

	value := 2

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
	// Value after mutation 1
	// ---
	// Value before mutation 5
	// Value of argument 1
	// Value after mutation 1
}

func swapArray(arg [2]int) {
	arg[0], arg[1] = arg[1], arg[0]
}

func Example_mutateArrayFunc() {
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

func Example_mustateSliceFunc() {
	slice := []int{1, 2}
	fmt.Printf("Slice before mutation %v\n", slice)
	swapSlice(slice)
	fmt.Printf("Slice after mutation %v\n", slice)

	// Output:
	// Slice before mutation [1 2]
	// Slice after mutation [2 1]
}

type exStruct struct {
	field1 int
}

func (e exStruct) getField1() int {
	return e.field1
}

func (e *exStruct) setField1(val int) {
	e.field1 = val
}

func Example_funcTypes() {
	// Function as value
	fnVar := func(i int) int { return i * 2 }
	fmt.Printf("Function as value %v\n", fnVar(3))

	// Output:
	// Function as value 6
}

func Example_funcMethods() {
	ex := &exStruct{field1: 1}
	var (
		field1FromMethod func() int    = ex.getField1
		field2SetMethod  func(arg int) = ex.setField1
	)
	fmt.Printf("Method getField1 %v\n", field1FromMethod())
	field2SetMethod(2)
	fmt.Printf("Method setField1 %v\n", ex.field1)

	// Output:
	// Method getField1 1
	// Method setField1 2
}
