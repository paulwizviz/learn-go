package encodejson

import (
	"encoding/json"
	"fmt"
)

type Ex1 struct {
	Field1 string
	Field2 string
}

func Example_marshallAllFields() {
	o := Ex1{
		Field1: "abc",
		Field2: "efg",
	}
	b, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	// Output:
	// {"Field1":"abc","Field2":"efg"}
}

func Example_unmarshallAllFields() {
	b := []byte(`{"Field1":"abc","Field2":"efg"}`)

	var n Ex1
	err := json.Unmarshal(b, &n)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

	// Output:
	// {abc efg}
}

func Example_marshallSomeFields() {
	o := Ex1{
		Field1: "abc",
	}
	b, err := json.Marshal(o)
	if err != nil {

	}
	fmt.Println(string(b))

	// Output:
	// {"Field1":"abc","Field2":""}
}

func Example_ummarshalSomeFields() {
	b := []byte(`{"Field1":"abc"}`)

	var n Ex1
	err := json.Unmarshal(b, &n)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

	b = []byte(`{"Field1":"abc", "Field2":""}`)
	err = json.Unmarshal(b, &n)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

	// Output:
	// {abc }
	// {abc }
}
