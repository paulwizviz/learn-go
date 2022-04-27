package inspect

import (
	"fmt"
	"reflect"
)

// Model
type IPerson interface{}

type Person struct {
	FirstName string `json:"firstname" c:"first_name"`
	Surname   string `json:"surname" c:"surname"`
}

type Address struct {
	Street  string
	City    string
	Country string
}

type PersonAddress struct {
	Person
	Address
	Landsize int
}

func fname() {}

// This example demonstrates the TypeOf function in the reflect package
func Example_typesOf() {

	intValue := 1
	typ := reflect.TypeOf(intValue)
	fmt.Printf("Primitive. Type: %T Value: %v Kind: %v\n", typ, typ, typ.Kind())

	intPtr := &intValue
	typ = reflect.TypeOf(intPtr)
	fmt.Printf("Pointer. Type: %T Value: %v Kind: %v Element: %v\n", typ, typ, typ.Kind(), typ.Elem())

	type Integer int
	var iVal Integer = Integer(1)
	typ = reflect.TypeOf(iVal)
	fmt.Printf("Alias. Type: %T Value: %v Kind: %v\n", typ, typ, typ.Kind())

	p := Person{
		FirstName: "firstname",
	}
	typ = reflect.TypeOf(p)
	fmt.Printf("Struct. Type: %T Value: %v Kind: %v\n", typ, typ, typ.Kind())

	pPtr := &Person{}
	typ = reflect.TypeOf(pPtr)
	fmt.Printf("Struct pointer. Type: %T Value: %v Kind: %v Element: %v\n", typ, typ, typ.Kind(), typ.Elem())

	var pI IPerson = &Person{}
	typ = reflect.TypeOf(pI)
	fmt.Printf("Struct interface. Type: %T Value: %v Kind: %v Element: %v\n", typ, typ, typ.Kind(), typ.Elem())

	typ = reflect.TypeOf(fname)
	fmt.Printf("Func. Type: %T Value: %v Kind: %v\n", typ, typ, typ.Kind())

	typ = reflect.TypeOf(func() string { return "" })
	fmt.Printf("Func. Type: %T Value: %v Kind: %v\n", typ, typ, typ.Kind())

	// Output:
	// Primitive. Type: *reflect.rtype Value: int Kind: int
	// Pointer. Type: *reflect.rtype Value: *int Kind: ptr Element: int
	// Alias. Type: *reflect.rtype Value: inspect.Integer Kind: int
	// Struct. Type: *reflect.rtype Value: inspect.Person Kind: struct
	// Struct pointer. Type: *reflect.rtype Value: *inspect.Person Kind: ptr Element: inspect.Person
	// Struct interface. Type: *reflect.rtype Value: *inspect.Person Kind: ptr Element: inspect.Person
	// Func. Type: *reflect.rtype Value: func() Kind: func
	// Func. Type: *reflect.rtype Value: func() string Kind: func
}

// This example demonstrates the use of ValueOf method in the reflect package
func Example_valueOf() {

	intValue := 1
	value := reflect.ValueOf(intValue)
	fmt.Printf("Primitive. Type: %T Value: %v Kind: %v\n", value, value, value.Kind())

	intPtr := &intValue
	value = reflect.ValueOf(intPtr)
	fmt.Printf("Primitive. Type: %T Value: %v Kind: %v Element: %v\n", value, "0x...", value.Kind(), value.Elem()) // Value == pointer addree

	p := Person{}
	value = reflect.ValueOf(p)
	fmt.Printf("Empty struct. Type: %T Value: %v Kind: %v\n", value, value, value.Kind())

	p1 := Person{
		FirstName: "firstname",
	}
	value = reflect.ValueOf(p1)
	fmt.Printf("Filled Struct. Type: %T Value: %v Kind: %v\n", value, value, value.Kind())

	ptr := &Person{
		FirstName: "firstname",
	}
	value = reflect.ValueOf(ptr)
	fmt.Printf("Filled Struct ptr. Type: %T Value: %v Kind: %v Element: %v\n", value, value, value.Kind(), value.Elem())

	value = reflect.ValueOf(fname)
	fmt.Printf("Func. Type: %T Value: %v Kind: %v\n", value, "0x...", value.Kind()) // value == address of func

	value = reflect.ValueOf(func() string { return "" })
	fmt.Printf("Func. Type: %T Value: %v Kind: %v\n", value, "0x...", value.Kind())

	// Output:
	// Primitive. Type: reflect.Value Value: 1 Kind: int
	// Primitive. Type: reflect.Value Value: 0x... Kind: ptr Element: 1
	// Empty struct. Type: reflect.Value Value: { } Kind: struct
	// Filled Struct. Type: reflect.Value Value: {firstname } Kind: struct
	// Filled Struct ptr. Type: reflect.Value Value: &{firstname } Kind: ptr Element: {firstname }
	// Func. Type: reflect.Value Value: 0x... Kind: func
	// Func. Type: reflect.Value Value: 0x... Kind: func
}

// Reflect indirect demonstrate the extraction of value
// of either a pointer or non-pointer
func Example_reflectIndirect() {
	p := Person{}
	value := reflect.Indirect(reflect.ValueOf(p))
	fmt.Printf("Type: %T Value: %v Kind: %v\n", value, value, value.Kind())

	ptr := &Person{}
	value = reflect.Indirect(reflect.ValueOf(ptr))
	fmt.Printf("Type: %T Value: %v Kind: %v\n", value, value, value.Kind())

	// Output:
	// Type: reflect.Value Value: { } Kind: struct
	// Type: reflect.Value Value: { } Kind: struct
}

func Example_fields() {
	pa := PersonAddress{}
	ptrPA := &pa

	fmt.Println("-- Using typeOf --")
	typ := reflect.TypeOf(pa)
	for index := 0; index < typ.NumField(); index++ {
		fmt.Println(typ.Field(index))
	}

	fmt.Println("---")

	typ = reflect.TypeOf(ptrPA)
	for index := 0; index < typ.Elem().NumField(); index++ {
		fmt.Println(typ.Elem().Field(index))
	}

	fmt.Println("-- Using valueOf --")
	value := reflect.Indirect(reflect.ValueOf(pa))
	for index := 0; index < value.NumField(); index++ {
		fmt.Println(value.Type().Field(index))
	}

	fmt.Println("---")

	value = reflect.Indirect(reflect.ValueOf(ptrPA))
	for index := 0; index < value.NumField(); index++ {
		fmt.Println(value.Type().Field(index))
	}

	// Output:
	// -- Using typeOf --
	// {Person  inspect.Person  0 [0] true}
	// {Address  inspect.Address  32 [1] true}
	// {Landsize  int  80 [2] false}
	// ---
	// {Person  inspect.Person  0 [0] true}
	// {Address  inspect.Address  32 [1] true}
	// {Landsize  int  80 [2] false}
	// -- Using valueOf --
	// {Person  inspect.Person  0 [0] true}
	// {Address  inspect.Address  32 [1] true}
	// {Landsize  int  80 [2] false}
	// ---
	// {Person  inspect.Person  0 [0] true}
	// {Address  inspect.Address  32 [1] true}
	// {Landsize  int  80 [2] false}
}

func Example_extractTags() {
	p := Person{}
	value := reflect.Indirect(reflect.ValueOf(p))
	tags := make(map[string]reflect.StructTag)
	for index := 0; index < value.Type().NumField(); index++ {
		fn := value.Type().Field(index).Name
		tag := value.Type().Field(index).Tag
		tags[fn] = tag
	}
	fmt.Println(tags)

	fmt.Println("----")

	ptr := &p
	value = reflect.Indirect(reflect.ValueOf(ptr))
	tags = make(map[string]reflect.StructTag)
	for index := 0; index < value.Type().NumField(); index++ {
		fn := value.Type().Field(index).Name
		tag := value.Type().Field(index).Tag
		tags[fn] = tag
	}
	fmt.Println(tags)

	// Output:
	// map[FirstName:json:"firstname" c:"first_name" Surname:json:"surname" c:"surname"]
	// ----
	// map[FirstName:json:"firstname" c:"first_name" Surname:json:"surname" c:"surname"]
}

func Example_inspectStructs() {
	p := PersonAddress{}
	result, err := ExtractStructMeta(p)
	fmt.Println(result, err)

	// Output:
	// &{inspect.PersonAddress [{Person  0x110aea0  0 [0] true} {Address  0x110eee0  32 [1] true} {Landsize  0x10feec0  80 [2] false}]} <nil>
}
