package reflectutil

import (
	"fmt"
	"reflect"
)

func ExtractTags(data interface{}) map[string]string {

	elm := reflect.ValueOf(data).Elem()
	for index := 0; index < elm.NumField(); index++ {
		tag := elm.Type().Field(index).Tag
		fmt.Println(elm, tag)
	}

	return nil
}

func ExtractFieldNames(data interface{}) {
	elm := reflect.ValueOf(data).Elem()
	for index := 0; index < elm.NumField(); index++ {
		fn := elm.Type().Field(index).Name
		ft := elm.Type().Field(index).Type.Name()
		fmt.Printf("Field name: %v, Field Type: %v", fn, ft)
	}
}
