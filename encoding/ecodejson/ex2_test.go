package encodejson

import (
	"encoding/json"
	"fmt"
)

type Ex21 struct {
	Field1 string  `json:"field1"`
	Field2 int     `json:"field2"`
	Field3 float32 `json:"field3"`
	Field4 string  `json:"-"`
}

func Example_marshalCorrectTags() {
	e := Ex21{
		Field1: "abc",
		Field2: 1,
		Field3: 1.4,
		Field4: "efg",
	}

	b, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))

	// Output:
	// {"field1":"abc","field2":1,"field3":1.4}
}

func Example_unmarshalCorrectTags() {
	b := []byte(`{"field1":"abc","field2":1,"field3":1.4}`)
	var ex2 Ex21
	err := json.Unmarshal(b, &ex2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ex2)

	// Output:
	// {abc 1 1.4 }
}

func Example_unmarshalJSONExtraField() {
	b := []byte(`{"field1":"abc","field2":1,"field3":1.4, "field4":"efg"}`)
	var ex2 Ex21
	err := json.Unmarshal(b, &ex2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ex2)

	// Output:
	// {abc 1 1.4 }
}

func Example_unmarshalJSONIncorrectFields() {
	b := []byte(`{"field23":"abc"`)
	var ex2 Ex21
	err := json.Unmarshal(b, &ex2)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// unexpected end of JSON input
}
