package main

import (
	"fmt"
	"reflect"
)

func Example_struct() {

	p := Person{
		FirstName: "Alice",
		Surname:   "Doe",
	}

	v := reflect.ValueOf(p)
	for i := 0; i < v.NumField(); i++ {
		sn := v.Type()
		fd := v.Field(i)
		tag := v.Type().Field(i).Tag
		fmt.Printf("Struct: %v Field: %v Tag: %v\n", sn, fd, tag)
	}

	// output:
	// Struct: main.Person Field: Alice Tag: json:"firstname" c:"first_name"
	// Struct: main.Person Field: Doe Tag: json:"surname" c:"surname"

}
