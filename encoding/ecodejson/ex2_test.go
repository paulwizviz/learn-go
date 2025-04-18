package encodejson

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Ex21 struct {
	Field1 string  `json:"field1"`
	Field2 int     `json:"field2"`
	Field3 float32 `json:"field3"`
	Field4 string  `json:"-"`
	Field5 string  `json:"field5,omitempty"`
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

type Ex22 struct {
	Field1 string  `json:"field1"`
	Field2 int     `json:"field2"`
	Field3 float32 `json:"field3"`
}

func Example_marshalEx22() {

	ex22 := Ex22{
		Field1: "abc",
		Field2: 1,
		Field3: 1.4,
	}

	type Alias Ex22
	aux := struct {
		*Alias
		Field3 string `json:"field3"`
	}{
		Alias:  (*Alias)(&ex22),
		Field3: fmt.Sprintf("%.2f", ex22.Field3),
	}

	b, err := json.Marshal(aux)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	// Output:
	// {"field1":"abc","field2":1,"field3":"1.40"}

}

func Example_unmarshalEx22() {
	ex22 := &Ex22{}

	type Alias Ex22
	aux := struct {
		*Alias
		Field3 string `json:"field3"`
	}{
		Alias: (*Alias)(ex22),
	}

	data := []byte(`{"field1":"abc","field2":1,"field3":"2.4"}`)
	err := json.Unmarshal(data, &aux)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(aux.Alias, aux.Field3)
	fmt.Println(ex22)
	f, err := strconv.ParseFloat(aux.Field3, 32)
	if err != nil {
		fmt.Println(err)
	}
	ex22.Field3 = float32(f)
	fmt.Println(ex22)

	// Output:
	// &{abc 1 0} 2.4
	// &{abc 1 0}
	// &{abc 1 2.4}
}
