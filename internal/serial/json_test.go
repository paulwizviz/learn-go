package serial

import (
	"encoding/json"
	"fmt"
	"time"
)

// Model
type task struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"` // Marshal this to int (Day) only
}

func (t *task) MarshalJSON() ([]byte, error) {
	type Alias task
	return json.Marshal(&struct {
		*Alias
		Created int64 `json:"created"` // Changing created from time.Time to Unix time
	}{
		Alias:   (*Alias)(t),
		Created: t.Created.Unix(),
	})
}

func (t *task) UnmarshalJSON(data []byte) error {
	type Alias task
	aux := &struct {
		*Alias
		Created int64 `json:"created"`
	}{
		Alias: (*Alias)(t),
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	t.Created = time.Unix(aux.Created, 0)
	return nil
}

func Example_marshalTask() {

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	tsk := &task{
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
	var actual task
	json.Unmarshal([]byte(s), &actual)
	fmt.Printf("Marshalled: %v Unmarshalled: %v", s, actual)

	// Output:
	// Marshalled: {"id":1,"name":"hello","created":1650297458} Unmarshalled: {1 hello 2022-04-18 16:57:38 +0100 BST}
}

type project interface {
	Identifier() string
}

type defaultProject struct {
	ID string `json:"id"`
}

func NewProject(id string) project {
	return &defaultProject{
		ID: id,
	}
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

func Example_marshalProject() {
	p := NewProject("hello")
	d, _ := json.Marshal(p)
	fmt.Println(string(d))

	// Output:
	// {"id":"hello"}
}

func Example_unmarshalProject() {
	p := NewProject("")
	// var p project -- It is not possible to unmarshal to p here
	err := json.Unmarshal([]byte(`{"id":"hello"}`), &p)
	fmt.Println(p, err)

	// Output:
	// &{hello} <nil>
}
