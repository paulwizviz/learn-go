package serial

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
)

type person struct {
	Firstname string
	Surname   string
}

func Example_marshalCBOR() {

	input := person{
		Firstname: "john",
		Surname:   "doe",
	}

	b, _ := cbor.Marshal(input)
	fmt.Printf("Input: %v Marshal string: %s Marshal bytes: %v", input, string(b), b)

	// Output:
}

func Example_umarshalCBOR() {

	input := person{
		Firstname: "john",
		Surname:   "doe",
	}

	b, _ := cbor.Marshal(input)

	var p person
	cbor.Unmarshal(b, &p)
	fmt.Println(p)

	// Output:
	// {john doe}
}
