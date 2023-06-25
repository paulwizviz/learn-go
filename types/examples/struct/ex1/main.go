package main

import "fmt"

type SimpleStruct struct {
	ID   int64
	Name string
}

func (s SimpleStruct) StructName() string {
	return s.Name
}

type CompoundStruct struct {
	SimpleStruct
	Name string
}

type AnotherCompoundStruct struct {
	SimpleStruct
	Name string
}

func (a AnotherCompoundStruct) StructName() string {
	return a.Name
}

func main() {
	e1 := CompoundStruct{
		SimpleStruct{
			ID:   123,
			Name: "Embedded",
		},
		"folks",
	}

	e2 := CompoundStruct{
		SimpleStruct: SimpleStruct{
			ID:   234,
			Name: "Embedded struct",
		},
		Name: "Hello",
	}

	fmt.Println(e1) // Output: {{123 Embedded} folks}
	fmt.Println(e2) // Output: {{234 Embedded struct} Hello}

	e3 := CompoundStruct{
		SimpleStruct: SimpleStruct{
			ID:   234,
			Name: "Embedded struct",
		},
		Name: "Hello",
	}

	fmt.Println(e3.ID)                // Output: 234
	fmt.Println(e3.Name)              // Output: Hello
	fmt.Println(e3.SimpleStruct.Name) // This disembiguate embedded struct name. Output: Embedded struct
	fmt.Println(e3.StructName())      // This calls the embedded struct method. Output: Embedded struct

	e4 := AnotherCompoundStruct{
		SimpleStruct: SimpleStruct{
			ID:   234,
			Name: "Embedded struct",
		},
		Name: "Hello",
	}

	fmt.Println(e4.StructName())              // This will always pick the method of the compounded type. Output: Hello
	fmt.Println(e4.SimpleStruct.StructName()) // This ensure that the embedded type method is called. Output: Embedded struct
}
