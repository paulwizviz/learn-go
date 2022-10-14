package embed_test

import (
	"fmt"

	"github.com/paulwizviz/learn-go/example/types/embed"
)

func Example_embeddedStruct() {

	e1 := embed.CompoundStruct{
		embed.SimpleStruct{
			ID:   123,
			Name: "Embedded",
		},
		"folks",
	}

	e2 := embed.CompoundStruct{
		SimpleStruct: embed.SimpleStruct{
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
	e := embed.CompoundStruct{
		SimpleStruct: embed.SimpleStruct{
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
