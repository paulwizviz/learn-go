package reflection

import (
	"fmt"
	"reflect"
)

func Example_inspectPerson() {

	p := Person{
		FirstName: "John",
		Surname:   "Doe",
	}
	v := reflect.ValueOf(p)
	t := reflect.TypeOf(p)
	for i := range v.NumField() {
		fmt.Printf("Field value: %v, Field name: %v, Exported: %v, Tag: %v\n", v.Field(i), t.Field(i).Name, t.Field(i).IsExported(), t.Field(i).Tag)
	}

	// Output:
	// Field value: John, Field name: FirstName, Exported: true, Tag: json:"firstname" c:"first_name"
	// Field value: Doe, Field name: Surname, Exported: true, Tag: json:"surname" c:"surname"
}
