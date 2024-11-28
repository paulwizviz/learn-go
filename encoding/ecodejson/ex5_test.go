package encodejson

import (
	"encoding/json"
	"fmt"
)

type Data interface{}

type Ex5 struct {
	ID   int  `json:"id"`
	Data Data `json:"data"`
}

type Data1 struct {
	Surname string `json:"surname"`
}

type Data2 struct {
	Firstname string `json:"firstname"`
}

func (e *Ex5) UnmarshalJSON(data []byte) error {
	type Alias Ex5
	temp := struct {
		Data json.RawMessage `json:"data"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	var data1 Data1
	if err := json.Unmarshal(temp.Data, &data1); err == nil && data1.Surname != "" {
		e.Data = data1
		return nil
	}

	var data2 Data2
	if err := json.Unmarshal(temp.Data, &data2); err == nil && data2.Firstname != "" {
		e.Data = data2
		return nil
	}
	return fmt.Errorf("unknown types")
}

func Example_marshalEx5() {
	data1 := Data1{
		Surname: "Doe",
	}
	ex5 := Ex5{
		ID:   1,
		Data: data1,
	}

	b, err := json.Marshal(ex5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))

	data2 := Data2{
		Firstname: "John",
	}
	ex5.Data = data2

	b, err = json.Marshal(ex5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))

	// Output:
	// {"id":1,"data":{"surname":"Doe"}}
	// {"id":1,"data":{"firstname":"John"}}
}

func Example_unmarshalData1() {
	b := []byte(` {"id":1,"data":{"surname":"Doe"}}`)
	var ex5 Ex5
	err := json.Unmarshal(b, &ex5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ex5)

	// Output:
	// {1 {Doe}}
}

func Example_unmarshalData2() {
	b := []byte(` {"id":1,"data":{"firstname":"John"}}`)
	var ex5 Ex5
	err := json.Unmarshal(b, &ex5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ex5)

	// Output:
	// {1 {John}}
}

func Example_unmarshalInvalidData() {
	b := []byte(` {"id":1,"data":{"name":"John"}}`)
	var ex5 Ex5
	err := json.Unmarshal(b, &ex5)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// unknown types
}
