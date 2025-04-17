package structs

import "fmt"

type Dog struct {
	name string
}

func (d *Dog) SetName(n string) {
	d.name = n
}

func (d Dog) Name() string {
	return d.name
}

func Example_dog() {
	d1 := Dog{
		name: "bob",
	}
	fmt.Println(d1.Name())
	d1.SetName("b")
	fmt.Println(d1.Name())

	d2 := &Dog{
		name: "skippy",
	}
	fmt.Println(d2.Name())
	d2.SetName("s")
	fmt.Println(d2.Name())

	// Output:
	// bob
	// b
	// skippy
	// s
}

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

func Example_cat() {
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
