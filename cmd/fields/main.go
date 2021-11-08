package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	Street  string
	City    string
	Country string
}

type Person struct {
	FirstName string
	Surname   string
}

type PersonAddress struct {
	Person
	Address
	Height int
}

type Addresses struct {
	Primary   Address
	Secondary Address
}

func ExtractFieldNames(data interface{}) {
	elm := reflect.ValueOf(data).Elem()

	for index := 0; index < elm.NumField(); index++ {
		fn := elm.Type().Field(index).Name
		ft := elm.Type().Field(index).Type

		if ft.Kind() == reflect.Struct {
			fmt.Printf("Field name: %v, Field Type: %v Kind: %v\n", fn, ft, ft.Kind())
			fd := elm.Type().Field(index)
			for i := 0; i < fd.Type.NumField(); i++ {
				sfn := fd.Type.Field(i).Name
				sft := fd.Type.Field(i).Type
				fmt.Printf("  Field name: %v Field Type: %v\n", sfn, sft)
			}
		} else {
			fmt.Printf("Field name: %v, Field Type: %v Kind: %v\n", fn, ft, ft.Kind())
		}
	}
}

func main() {
	fmt.Println("=====================")
	fmt.Println("Reflection of struct types")
	fmt.Println("=====================")
	fmt.Println("")
	fmt.Println("======== Address ==========")
	a := &Address{
		Street:  "10th Street",
		City:    "Some city",
		Country: "Some country",
	}
	ExtractFieldNames(a)

	fmt.Println("======== PersonAddress ==========")
	pa := &PersonAddress{
		Person{
			FirstName: "John",
			Surname:   "Doe",
		},
		Address{
			Street:  "10 Street",
			City:    "Some city",
			Country: "Country",
		},
		180,
	}
	ExtractFieldNames(pa)

	fmt.Println("======== Addresses ===========")
	addr := &Addresses{
		Primary: Address{
			Street:  "11 Street",
			City:    "A city",
			Country: "A country",
		},
		Secondary: Address{
			Street:  "12 Street",
			City:    "Another city",
			Country: "Another country",
		},
	}
	ExtractFieldNames(addr)
}
