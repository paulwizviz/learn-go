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
	now := time.Now()

	tsk := &task{
		ID:      1,
		Name:    "hello",
		Created: now,
	}

	b, _ := json.Marshal(tsk)
	fmt.Printf("Input struct: %v | Marshalled: %v", tsk, string(b))

	// Output:
	// Input struct: &{1 hello 2022-04-27 11:12:14.629277 +0100 BST m=+0.001260505} Marshalled: {"id":1,"name":"hello","created":1651054334}
}

func Example_unmarshalTask() {
	s := `{"id":1,"name":"hello","created":1650297458}`
	var actual task
	json.Unmarshal([]byte(s), &actual)
	fmt.Printf("Marshalled: %v Unmarshalled: %v", s, actual)

	// Output:
	// Marshalled: {"id":1,"name":"hello","created":1650297458} Unmarshalled: {1 hello 2022-04-18 16:57:38 +0100 BST}
}
