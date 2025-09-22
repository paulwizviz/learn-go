package reflection

import (
	"fmt"
	"reflect"
)

func Example_primitveTypeOf() {
	// int type
	intValue := 1
	t := reflect.TypeOf(intValue)
	fmt.Printf("Primitive. Type: %v Kind: %v\n", t, t.Kind())

	// int pointer
	intPtr := &intValue
	t = reflect.TypeOf(intPtr)
	fmt.Printf("Pointer. Type: %v Kind: %v Element: %v\n", t, t.Kind(), t.Elem())

	// nil pointer
	intPtr = nil
	t = reflect.TypeOf(intPtr)
	fmt.Printf("Pointer. Type: %v Kind: %v Element: %v\n", t, t.Kind(), t.Elem())

	// Type wrapper
	type Integer int
	var iVal Integer = Integer(1)
	t = reflect.TypeOf(iVal)
	fmt.Printf("Alias. Type: %v Kind: %v\n", t, t.Kind())

	// output:
	// Primitive. Type: int Kind: int
	// Pointer. Type: *int Kind: ptr Element: int
	// Pointer. Type: *int Kind: ptr Element: int
	// Alias. Type: reflection.Integer Kind: int
}

func Example_collectionTypeOf() {

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

func Example_func() {

	t := reflect.TypeOf(fn)
	fmt.Printf("Name function. Type: %v Kind: %v\n", t, t.Kind())

	var fn1 fname
	t = reflect.TypeOf(fn1)
	fmt.Printf("Named func. Type: %v Kind: %v\n", t, t.Kind())

	// output:
	// Name function. Type: func() Kind: func
	// Named func. Type: reflection.fname Kind: func
}

func Example_interfaceTypeOf() {
	var h1 Human = Person{}
	t := reflect.TypeOf(h1)
	fmt.Printf("Struct interface. Type: %v Kind: %v\n", t, t.Kind())

	var h2 Human = &Person{}
	t = reflect.TypeOf(h2)
	fmt.Printf("Struct pointer interface. Type: %v Kind: %v Element: %v\n", t, t.Kind(), t.Elem())

	// Output:
	// Struct interface. Type: reflection.Person Kind: struct
	// Struct pointer interface. Type: *reflection.Person Kind: ptr Element: reflection.Person

}

func Example_structTypeOf() {

	p := Person{}
	t := reflect.TypeOf(p)
	fmt.Printf("Struct. Type: %v Kind: %v NumField: %v\n", t, t.Kind(), t.NumField())

	pPtr := &Person{}
	t = reflect.TypeOf(pPtr)
	fmt.Printf("Struct pointer. Type: %v Kind: %v Element: %v\n", t, t.Kind(), t.Elem())

	// output:
	// Struct. Type: reflection.Person Kind: struct NumField: 2
	// Struct pointer. Type: *reflection.Person Kind: ptr Element: reflection.Person
}
