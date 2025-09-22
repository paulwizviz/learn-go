package reflection

import (
	"fmt"
	"reflect"
)

func Example_indirectPerson() {

	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	p := Person{
		FirstName: "John",
		Surname:   "Doe",
	}

	v := reflect.ValueOf(p)
	fmt.Printf("NumField: %v\n", v.NumField())

	indirect := reflect.Indirect(reflect.ValueOf(&p))
	fmt.Printf("NumField: %v\n", indirect.NumField())

	v = reflect.ValueOf(&p)
	v.NumField() // Trigger error

	// Output:
	// NumField: 2
	// NumField: 2
	// reflect: call of reflect.Value.NumField on ptr Value
}

func Example_indirectHuman() {

	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	var h1 Human = Person{}
	v := reflect.ValueOf(h1)
	fmt.Printf("NumFields: %v\n", v.NumField())

	var h2 Human = &Person{}
	v = reflect.ValueOf(h2)

	indirect := reflect.Indirect(v)
	fmt.Printf("NumFields: %v\n", indirect.NumField())
	v.NumField() // Error

	// Output:
	// NumFields: 2
	// NumFields: 2
	// reflect: call of reflect.Value.NumField on ptr Value
}

func Example_catValueOf() {
	c := Cat{
		name: "Meow",
	}
	v := reflect.ValueOf(c)
	t := reflect.TypeOf(&c)
	fmt.Println(v.NumMethod())
	fmt.Println(t.NumMethod())
	// Output:
}

func Example_catIndirect() {

	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	c := &Cat{
		name: "Tabby",
	}

	var a1 Animal = c
	v := reflect.ValueOf(a1)
	ind := reflect.Indirect(v) // Getting the value of underlying type
	fmt.Printf("NumFields: %v\n", ind.NumField())
	fmt.Printf("NumMethods: %v\n", ind.NumMethod()) // This will not return pointer receiver (*c) method
	fmt.Printf("NumMethods: %v\n", v.NumMethod())
	v.NumField() // Unable to retrieve number of fields

	// Output:
	// NumFields: 1
	// NumMethods: 1
	// NumMethods: 2
	// reflect: call of reflect.Value.NumField on ptr Value
}
