package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	FirstName string `json:"firstname" c:"first_name"`
	Surname   string `json:"surname" c:"surname"`
}

func ExtractTags(data interface{}) {
	elm := reflect.ValueOf(data).Elem()
	for index := 0; index < elm.NumField(); index++ {
		fn := elm.Type().Field(index).Name
		tag := elm.Type().Field(index).Tag
		fmt.Println("Field name: ", fn)

		for _, tElm := range SplitTags(tag) {
			kv := strings.Split(tElm, ":")
			fmt.Printf("  tag: %v value: %v\n", tElm, tag.Get(kv[0]))
		}
	}
}

func SplitTags(tag reflect.StructTag) []string {
	elms := strings.Split(string(tag), " ")
	return elms
}

func main() {
	fmt.Println("====== Extract Person Tag ===========")
	p := &Person{}

	result, _ := json.Marshal(p)
	fmt.Println("->", string(result))
	ExtractTags(p)
}
