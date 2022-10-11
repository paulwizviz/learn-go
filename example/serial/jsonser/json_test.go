package jsonser_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/paulwizviz/learn-go/example/serial/jsonser"
)

// Examples
func Example_marshalTask() {

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	tsk := &jsonser.Task{
		ID:      1,
		Name:    "hello",
		Created: tm,
	}

	b, _ := json.Marshal(tsk)
	fmt.Printf("Input struct: %v | Marshalled: %v", tsk, string(b))

	// Output:
	// Input struct: &{1 hello 2006-01-02 15:04:05 +0000 UTC} | Marshalled: {"id":1,"name":"hello","created":1136214245}
}

func Example_unmarshalTask() {
	s := `{"id":1,"name":"hello","created":1650297458}`
	var actual Task
	json.Unmarshal([]byte(s), &actual)
	fmt.Printf("Marshalled: %v Unmarshalled: %v", s, actual)

	// Output:
	// Marshalled: {"id":1,"name":"hello","created":1650297458} Unmarshalled: {1 hello 2022-04-18 16:57:38 +0100 BST}
}

func Example_marshalProject() {
	p := jsonser.NewProject("hello")
	d, _ := json.Marshal(p)
	fmt.Println(string(d))

	// Output:
	// {"id":"hello"}
}

func Example_unmarshalProject() {
	p := jsonser.NewProject("")
	// var p project -- It is not possible to unmarshal to p here
	err := json.Unmarshal([]byte(`{"id":"hello"}`), &p)
	fmt.Println(p, err)

	// Output:
	// &{hello} <nil>
}
