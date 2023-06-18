package embed_test

import (
	"fmt"
)

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

func Example_embeddedStruct() {

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

func Example_embeddedAttr() {
	e := CompoundStruct{
		SimpleStruct: SimpleStruct{
			ID:   234,
			Name: "Embedded struct",
		},
		Name: "Hello",
	}

	fmt.Println(e.ID)
	fmt.Println(e.Name)
	fmt.Println(e.SimpleStruct.Name) // This disembiguate embedded struct name
	fmt.Println(e.StructName())      // Get struct name

	// Output:
	// 234
	// Hello
	// Embedded struct
	// Embedded struct
}
