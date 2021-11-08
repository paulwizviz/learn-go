package reflectutil

import (
	"fmt"
	"reflect"
)

func ExtractTags(data interface{}) {
	elm := reflect.ValueOf(data).Elem()
	for index := 0; index < elm.NumField(); index++ {
		tag := elm.Type().Field(index).Tag
		fmt.Println(elm, tag)
	}
}
