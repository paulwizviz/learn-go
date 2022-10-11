package inspect

// Model
type IPerson interface{}

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

func fname() {}
