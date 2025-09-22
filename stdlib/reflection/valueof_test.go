package reflection

import (
	"fmt"
	"reflect"
)

func Example_primitiveValueOf() {

	// int type
	intValue := 10
	v := reflect.ValueOf(intValue)
	fmt.Printf("Primitive. Value: %v Kind: %v Type: %v\n", v, v.Kind(), v.Type())

	// int pointer
	intPtr := &intValue
	v = reflect.ValueOf(intPtr)
	fmt.Printf("Pointer. Value: <v is pointer address> Kind: %v Type: %v Element: %v\n", v.Kind(), v.Type(), v.Elem())

	// nil pointer
	intPtr = nil
	v = reflect.ValueOf(intPtr)
	fmt.Printf("Nil pointer. Value: <v is pointer address> Kind: %v Type: %v Element: %v IsNil: %v\n", v.Kind(), v.Type(), v.Elem(), v.IsNil())

	// Type wrapper
	type Integer int
	var iVal Integer = Integer(1)
	v = reflect.ValueOf(iVal)
	fmt.Printf("Alias. Value: %v Kind: %v\n", v, v.Kind())

	// output:
	// Primitive. Value: 10 Kind: int Type: int
	// Pointer. Value: <v is pointer address> Kind: ptr Type: *int Element: 10
	// Nil pointer. Value: <v is pointer address> Kind: ptr Type: *int Element: <invalid reflect.Value> IsNil: true
	// Alias. Value: 1 Kind: int
}

func Example_valueOf() {
	var h1 Human = Person{
		FirstName: "John",
		Surname:   "Doe",
	}
	v := reflect.ValueOf(h1)
	fmt.Printf("Struct interface. Value: %v Kind: %v\n", v, v.Kind())

	var h2 Human = &Person{
		FirstName: "John",
		Surname:   "Doe",
	}
	v = reflect.ValueOf(h2)
	fmt.Printf("Struct interface. Value: %v Kind: %v\n", v, v.Kind())

	// Output:
	// Struct interface. Value: {John Doe} Kind: struct
	// Struct interface. Value: &{John Doe} Kind: ptr
}

func Example_collectionValueOf() {

	m := map[string]int{
		"a": 1,
		"c": 2,
	}
	v := reflect.ValueOf(m)
	fmt.Printf("Value: %v Type: %v Kind: %v Key: %v Element: %v\n", v, v.Type(), v.Kind(), v.Type().Key(), v.Type().Elem())

	s := []int{1, 2}
	v = reflect.ValueOf(s)
	fmt.Printf("Value: %v Type: %v Kind: %v Element: %v\n", v, v.Type(), v.Kind(), v.Type().Elem())

	// output:
	// Value: map[a:1 c:2] Type: map[string]int Kind: map Key: string Element: int
	// Value: [1 2] Type: []int Kind: slice Element: int

}

func Example_structValueOf() {

	// Struct
	p := Person{
		FirstName: "John",
	}
	v := reflect.ValueOf(p)
	fmt.Printf("Struct. Kind: %v Type: %v Value: %v\n", v.Kind(), v.Type(), v)

	// Pointer to struct
	v = reflect.ValueOf(&p)
	fmt.Printf("Struct pointer. Kind: %v Type: %v Value: %v\n", v.Kind(), v.Type(), v)

	// Set pointer to nil
	var pt *Person
	v = reflect.ValueOf(pt)
	fmt.Printf("nil. Value: %v Kind: %v\n", v, v.Kind())

	// output:
	// Struct. Kind: struct Type: main.Person Value: {John }
	// Struct pointer. Kind: ptr Type: *main.Person Value: &{John }
	// Value: <nil> Kind: ptr
}

func Example_funcValueOf() {

	v := reflect.ValueOf(fn)
	fmt.Printf("Type of v: %T\n", v)
	fmt.Printf("Value: <v is address of v> Type: %v Kind: %v\n", v.Type(), v.Kind())

	var fn1 fname
	v = reflect.ValueOf(fn1)
	fmt.Printf("Value: <v is address> Type: %v Kind: %v\n", v.Type(), v.Kind())

	fn1 = func(s string) error {
		return nil
	}
	v = reflect.ValueOf(fn1)
	fmt.Printf("Value: <v is address> Type: %v Kind: %v\n", v.Type(), v.Kind())

	// output:
	// Type of v: reflect.Value
	// Value: <v is address of v> Type: func() Kind: func
	// Value: <v is address> Type: reflection.fname Kind: func
	// Value: <v is address> Type: reflection.fname Kind: func
}
