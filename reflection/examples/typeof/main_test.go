package main

import (
	"fmt"
	"reflect"
)

func Example_primitve() {
	// int type
	intValue := 1
	typ := reflect.TypeOf(intValue)
	fmt.Printf("Primitive. Type: %v Kind: %v\n", typ, typ.Kind())

	// int pointer
	intPtr := &intValue
	typ = reflect.TypeOf(intPtr)
	fmt.Printf("Pointer. Type: %v Kind: %v Element: %v\n", typ, typ.Kind(), typ.Elem())

	// nil pointer
	intPtr = nil
	typ = reflect.TypeOf(intPtr)
	fmt.Printf("Pointer. Type: %v Kind: %v Element: %v\n", typ, typ.Kind(), typ.Elem())

	// Type wrapper
	type Integer int
	var iVal Integer = Integer(1)
	typ = reflect.TypeOf(iVal)
	fmt.Printf("Alias. Type: %v Kind: %v\n", typ, typ.Kind())

	// output:
	// Primitive. Type: int Kind: int
	// Pointer. Type: *int Kind: ptr Element: int
	// Pointer. Type: *int Kind: ptr Element: int
	// Alias. Type: main.Integer Kind: int
}

func Example_collection() {

	m := map[string]int{
		"a": 1,
		"c": 2,
	}
	typ := reflect.TypeOf(m)
	fmt.Printf("Type: %v Kind: %v Key: %v Element: %v\n", typ, typ.Kind(), typ.Key(), typ.Elem())

	s := []int{1, 1}
	typ = reflect.TypeOf(s)
	fmt.Printf("Type: %v Kind: %v Element: %v\n", typ, typ.Kind(), typ.Elem())

	// output:
	// Type: map[string]int Kind: map Key: string Element: int
	// Type: []int Kind: slice Element: int

}

func Example_struct() {

	p := Person{}
	typ := reflect.TypeOf(p)
	fmt.Printf("Struct. Type: %v Kind: %v\n", typ, typ.Kind())

	pPtr := &Person{}
	typ = reflect.TypeOf(pPtr)
	fmt.Printf("Struct pointer. Type: %v Kind: %v Element: %v\n", typ, typ.Kind(), typ.Elem())

	var h Human = &Person{}
	typ = reflect.TypeOf(h)
	fmt.Printf("Struct interface. Type: %v Kind: %v Element: %v\n", typ, typ.Kind(), typ.Elem())

	// output:
	// Struct. Type: main.Person Kind: struct
	// Struct pointer. Type: *main.Person Kind: ptr Element: main.Person
	// Struct interface. Type: *main.Person Kind: ptr Element: main.Person
}

func Example_func() {

	typ := reflect.TypeOf(fn)
	fmt.Printf("Name function. Type: %v Kind: %v\n", typ, typ.Kind())

	var fn1 fname
	typ = reflect.TypeOf(fn1)
	fmt.Printf("Named func. Type: %v Kind: %v\n", typ, typ.Kind())

	// output:
	// Name function. Type: func() Kind: func
	// Named func. Type: main.fname Kind: func
}
