package structs

import "fmt"

type AStruct struct {
	field1 int
	field2 string
}

type BStruct struct {
	AStruct
}

type CStruct struct {
	*AStruct
}

func Example_equalByValue() {
	a1 := AStruct{
		field1: 1,
		field2: "hello",
	}

	a2 := AStruct{
		field1: 1,
		field2: "hello",
	}

	fmt.Println("Same: ", a1 == a2)

	// Output:
	// Same:  true
}

func Example_bStruct() {

	a1 := BStruct{
		AStruct{1, "hello"},
	}

	a2 := BStruct{
		AStruct{1, "hello"},
	}

	fmt.Printf("Same %v", a1 == a2)

	// Output:
	// Same true
}

func Example_cStruct() {

	a1 := CStruct{
		&AStruct{1, "hello"},
	}

	a2 := CStruct{
		&AStruct{1, "hello"},
	}

	fmt.Println(a1 == a2)

	// Output:
	// false
}

func Example_equalPtr() {

	a1 := &AStruct{
		field1: 1,
		field2: "hello",
	}

	a2 := a1
	fmt.Printf("Same ptrs: %v %v\n", a1, a2)

	a1.field1 = 4
	fmt.Printf("Same ptrs: %v %v\n", a1, a2)

	a3 := &AStruct{
		field1: 4,
		field2: "hello",
	}
	fmt.Printf("Different ptrs: %v %v", a1 != a3, *a1 == *a3)

	// Output:
	// Same ptr: &{1 hello} &{1 hello}
	// Same ptr: &{4 hello} &{4 hello}
	// Different ptrs: true true
}
