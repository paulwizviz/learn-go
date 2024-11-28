package encodejson

import (
	"encoding/json"
	"fmt"
	"time"
)

// Ex3
type Ex3 struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"` // Marshal this to int (Day) only
}

func (t *Ex3) MarshalJSON() ([]byte, error) {
	type Alias Ex3
	return json.Marshal(&struct {
		*Alias
		Created int64 `json:"created"` // Changing created from time.Time to Unix time
	}{
		Alias:   (*Alias)(t),
		Created: t.Created.Unix(),
	})
}

func (t *Ex3) UnmarshalJSON(data []byte) error {
	type Alias Ex3
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

// Examples
func Example_marshalTask() {

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	tsk := &Ex3{
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
	var actual Ex3
	json.Unmarshal([]byte(s), &actual)
	fmt.Printf("Marshalled: %v Unmarshalled: %v", s, actual)

	// Output:
	// Marshalled: {"id":1,"name":"hello","created":1650297458} Unmarshalled: {1 hello 2022-04-18 16:57:38 +0100 BST}
}
