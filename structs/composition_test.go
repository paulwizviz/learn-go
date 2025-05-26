package structs

import "fmt"

type SimpleStruct struct {
	ID   int64
	Name string
}

func Example_simpleStruct() {
	s := SimpleStruct{
		ID:   123,
		Name: "Hello",
	}

	fmt.Println(s)
	fmt.Println(s.ID)
	fmt.Println(s.Name)

	// Output:
	// {123 Hello}
	// 123
	// Hello
}

type CompoundStruct struct {
	SimpleStruct
	Name string
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
			Name: "Embedded_struct",
		},
		Name: "Hello",
	}

	fmt.Println(e1)
	fmt.Println(e2)

	// Output:
	// {{123 Embedded} folks}
	// {{234 Embedded_struct} Hello}
}

type AnotherCompoundStruct struct {
	*SimpleStruct
	Name string
}

func Example_anotherCompoundStruct() {
	s := SimpleStruct{
		ID:   123,
		Name: "Another_Compound",
	}

	e1 := AnotherCompoundStruct{
		SimpleStruct: &s,
		Name:         "folks",
	}

	fmt.Println(e1.ID)
	fmt.Println(e1.Name)
	fmt.Println(e1.SimpleStruct.Name)

	// Output:
	// 123
	// folks
	// Another_Compound
}
