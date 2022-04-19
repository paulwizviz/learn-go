package serial

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

// Test cases
func TestMarshalTask(t *testing.T) {
	now := time.Now()
	expected := now.Unix()

	tsk := &task{
		ID:      1,
		Name:    "hello",
		Created: now,
	}

	b, _ := json.Marshal(tsk)
	// b = {"id":1,"name":"hello","created":1650297458}
	s := string(b)
	tm, _ := strconv.Atoi(s[33 : len(s)-1])
	assert.Equal(t, expected, int64(tm))
}

func TestUnmarshalTask(t *testing.T) {
	s := `{"id":1,"name":"hello","created":1650297458}`
	var actual task
	json.Unmarshal([]byte(s), &actual)
	expected := task{
		ID:      1,
		Name:    "hello",
		Created: time.Unix(1650297458, 0),
	}
	assert.Equal(t, expected, actual)
}
