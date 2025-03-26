package reflection

type Human interface{}

type Person struct {
	FirstName string `json:"firstname" c:"first_name"`
	Surname   string `json:"surname" c:"surname"`
}

type Address struct {
	Street  string
	City    string
	Country string
}

type PersonAddress struct {
	Person
	Address
	Landsize int
}

type Animal interface {
	Name() string
	SetName(name string)
}

type Cat struct {
	name string
}

func (c Cat) Name() string {
	return c.name
}

func (c *Cat) SetName(name string) {
	c.name = name
}

func fn() {}

type fname func(s string) error
