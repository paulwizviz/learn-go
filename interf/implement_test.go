package interf

import "fmt"

type Printer interface {
	Print()
}

// Number is an alias for int type
// with a method to inplement printer
// interface.
type Number int

func (n Number) Print() {
	fmt.Printf("Number value: %v\n", n)
}

// PrintableType is a struct with
type PrintableType struct {
	msg string
}

func (p PrintableType) Print() {
	fmt.Printf("PrintableType message: %v\n", p.msg)
}

// FuncType is a type of function implementing
// the print method of the printer
type FuncType func()

func (f FuncType) Print() {
	f()
}

// PrintServices is a function that consume a printer
// type
func PrintServices(p Printer) {
	p.Print()
}

func Example_interfaceImpt() {

	var n Number = 1_000
	PrintServices(n)

	p := PrintableType{
		msg: "Says hello",
	}
	PrintServices(p)

	var f FuncType = func() {
		fmt.Printf("Function type: says Hello")
	}
	PrintServices(f)

	// Output:
	// Number value: 1000
	// PrintableType message: Says hello
	// Function type: says Hello

}
