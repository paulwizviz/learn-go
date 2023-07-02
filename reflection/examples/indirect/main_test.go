package main

import (
	"fmt"
	"reflect"
)

func Example_indirect() {

	p := Person{
		FirstName: "John",
		Surname:   "Doe",
	}

	v := reflect.ValueOf(p)
	fmt.Printf("Struct. Number of fields: %v\n", v.NumField())

	// Pointer. You can't get number of fields from reflect.ValueOf
	id := reflect.Indirect(reflect.ValueOf(&p))
	fmt.Printf("Struct ptr. Number of fields: %v\n", id.NumField())

	pa := PersonAddress{}
	v = reflect.ValueOf(pa)
	fmt.Printf("Compound struct. Number of fields: %v\n", v.NumField())

	// output:
	// Struct. Number of fields: 2
	// Struct ptr. Number of fields: 2
	// Compound struct. Number of fields: 3
}
