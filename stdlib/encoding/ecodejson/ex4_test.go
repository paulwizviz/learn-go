package encodejson

import (
	"encoding/json"
	"fmt"
)

// Ex4 Interface
type Ex4 interface {
	Identifier() string
}

// Implementations of project type
type defaultProject struct {
	ID string `json:"id"`
}

func (d *defaultProject) Identifier() string {
	return d.ID
}

func (d *defaultProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID string `json:"id"`
	}{
		ID: d.ID,
	})
}

func (d *defaultProject) UnmarshalJSON(data []byte) error {
	var alt struct {
		ID string `json:"id"`
	}
	err := json.Unmarshal(data, &alt)
	if err != nil {
		return err
	}
	d.ID = alt.ID
	return nil
}

func NewProject(id string) Ex4 {
	return &defaultProject{
		ID: id,
	}
}

func Example_marshalEx4() {
	p := NewProject("hello")
	d, _ := json.Marshal(p)
	fmt.Println(string(d))

	// Output:
	// {"id":"hello"}
}

func Example_unmarshalEx4() {
	p := NewProject("")
	// var p project -- It is not possible to unmarshal to p here
	err := json.Unmarshal([]byte(`{"id":"hello"}`), &p)
	fmt.Println(p, err)

	// Output:
	// &{hello} <nil>
}
