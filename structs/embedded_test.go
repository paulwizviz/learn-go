package structs

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

func Example_compoundStruct() {
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

	fmt.Println(e1)
	fmt.Println(e2)

	// Output:
	// {{123 Embedded} folks}
	// {{234 Embedded struct} Hello}
}

func Example_accessField() {
	e3 := CompoundStruct{
		SimpleStruct: SimpleStruct{
			ID:   234,
			Name: "Embedded struct",
		},
		Name: "Hello",
	}

	fmt.Println(e3.ID)
	fmt.Println(e3.Name)
	fmt.Println(e3.SimpleStruct.Name) // This disembiguate embedded struct name.
	fmt.Println(e3.StructName())      // This calls the embedded struct method.

	e4 := AnotherCompoundStruct{
		SimpleStruct: SimpleStruct{
			ID:   234,
			Name: "Embedded struct",
		},
		Name: "Hello",
	}

	fmt.Println(e4.StructName())              // This will always pick the method of the compounded type.
	fmt.Println(e4.SimpleStruct.StructName()) // This ensure that the embedded type method is called.

	// Output:
	// 234
	// Hello
	// Embedded struct
	// Embedded struct
	// Hello
	// Embedded struct
}
