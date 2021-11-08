package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("=================")
	fmt.Println(" Reflection of primitive types ")
	fmt.Println("=================")
	reflectIntType()
	fmt.Println("----")
	reflectFloatType()
	fmt.Println("----")
	reflectBoolType()
}

func reflectIntType() {
	var intValue int = 100
	value := reflect.ValueOf(&intValue)
	fmt.Printf("Reflected pointer value: %v\n", value.Pointer())
	fmt.Printf("Reflected underlying value: %v\n", value.Elem())
	fmt.Printf("Reflected type value: %v\n", value.Type())
}

func reflectFloatType() {
	var fValue float32 = 100.01
	value := reflect.ValueOf(&fValue)
	fmt.Printf("Reflected pointer value: %v\n", value.Pointer())
	fmt.Printf("Reflected underlying value: %v\n", value.Elem())
	fmt.Printf("Reflected type value: %v\n", value.Type())
}

func reflectBoolType() {
	var bValue bool = true
	value := reflect.ValueOf(&bValue)
	fmt.Printf("Reflected pointer value: %v\n", value.Pointer())
	fmt.Printf("Reflected underlying value: %v\n", value.Elem())
	fmt.Printf("Reflected type value: %v\n", value.Type())
}
