package structs

import "fmt"

type Cat struct {
	name string
}

func (c Cat) Name() string {
	return c.name
}

// This is a syntatic sugar for this
// func SetName(c *Cat, n string)
func (c *Cat) SetName(n string) {
	c = nil // This is unlikely scenario but the compiler does not pick this up.
	c.name = n
}

func Example_failed() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	c := &Cat{
		name: "cat",
	}
	fmt.Println(c.Name())
	c.SetName("b")

	// Output:
	// cat
	// runtime error: invalid memory address or nil pointer dereference
}
