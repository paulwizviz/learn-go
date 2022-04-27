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

type human interface {
	Firstname() string
	Surname() string
}

type defaultHuman struct {
	firstname string
	surname   string
}

func (d *defaultHuman) Firstname() string {
	return d.firstname
}

func (d *defaultHuman) Surname() string {
	return d.surname
}

func (d *defaultHuman) MarshalCBOR() ([]byte, error) {
	return cbor.Marshal(&struct {
		Firstname string
		Surname   string
	}{
		Firstname: d.firstname,
		Surname:   d.surname,
	})
}

func (d *defaultHuman) UnmarshalCBOR(data []byte) error {
	var aux struct {
		Firstname string
		Surname   string
	}
	if err := cbor.Unmarshal(data, &aux); err != nil {
		return err
	}
	d.firstname = aux.Firstname
	d.surname = aux.Surname
	return nil
}

func NewHuman(firstname, surname string) human {
	return &defaultHuman{
		firstname: firstname,
		surname:   surname,
	}
}

func Example_marshalInterfaceCBOR() {
	h := NewHuman("john", "doe")
	b, _ := cbor.Marshal(h)
	for _, r := range b {
		fmt.Print(string(r))
	}

	// Output:
	// ¢a1djohna2cdoe
}

func Example_unmarshalInterfaceCBOR() {
	h := NewHuman("john", "doe")
	b, _ := cbor.Marshal(h)

	h1 := NewHuman("", "")
	cbor.Unmarshal(b, &h1)
	fmt.Println(h1)
	// Output:
	// &{john doe}
}
