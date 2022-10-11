package jsonser

import (
	"encoding/json"
	"time"
)

// Task
type Task struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"` // Marshal this to int (Day) only
}

func (t *Task) MarshalJSON() ([]byte, error) {
	type Alias Task
	return json.Marshal(&struct {
		*Alias
		Created int64 `json:"created"` // Changing created from time.Time to Unix time
	}{
		Alias:   (*Alias)(t),
		Created: t.Created.Unix(),
	})
}

func (t *Task) UnmarshalJSON(data []byte) error {
	type Alias Task
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

// Project Interface
type Project interface {
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

func NewProject(id string) Project {
	return &defaultProject{
		ID: id,
	}
}
