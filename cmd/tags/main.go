package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	FirstName string `json:"firstname"`
	Surname   string `json:"surname"`
}

func ExtractTags(data interface{}) {
	elm := reflect.ValueOf(data).Elem()
	for index := 0; index < elm.NumField(); index++ {
		fn := elm.Type().Field(index).Name
		tag := elm.Type().Field(index).Tag
		fmt.Printf("Field name: %v Tag: %v\n", fn, tag)
	}
}

func main() {
	fmt.Println("====== Extract Person Tag ===========")
	p := &Person{}
	ExtractTags(p)
}
