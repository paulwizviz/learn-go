package embed

import (
	"fmt"
)

type simpleStruct struct {
	ID   int64
	Name string
}

type embeddedStruct struct {
	simpleStruct
	name string
}

func Example_embeddedStruct() {

	e1 := embeddedStruct{
		simpleStruct{
			ID:   123,
			Name: "Embedded",
		},
		"folks",
	}

	e2 := embeddedStruct{
		simpleStruct: simpleStruct{
			ID:   234,
			Name: "Embedded struct",
		},
		name: "Hello",
	}

	fmt.Println(e1)
	fmt.Println(e2)

	// Output:
	// {{123 Embedded} folks}
	// {{234 Embedded struct} Hello}

}
