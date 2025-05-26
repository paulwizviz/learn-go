package structs

import "fmt"

// Number is an alias for float64
type Number float64

func (n Number) Increment() {
	n++ // This will not mutate original value
}

// Pointer type receiver to mutate underlying value
func (n *Number) IncreaseWith(num Number) {
	*n = *n + num //
}

func Example_number() {
	var n Number = 1.234
	fmt.Println("Original value: ", n)

	n.Increment()
	fmt.Println("After increment: ", n)

	n.IncreaseWith(4.0)
	fmt.Println("Increase by 4.0: ", n)

	var n1 *Number = &n
	n.IncreaseWith(10.00)
	fmt.Println("Increase by 10.00", *n1)

	// Output:
	// Original value:  1.234
	// After increment:  1.234
	// Increase by 4.0:  5.234
	// Increase by 10.00 15.234
}

// Human represents a person entity
type Human struct {
	name string
}

func (p Human) Name() string {
	return p.name
}

func (p *Human) SetName(name string) {
	p.name = name
}

func Example_human() {
	p1 := &Human{
		name: "bob",
	}

	fmt.Println(p1)

	p1.SetName("alice")
	fmt.Println(p1.Name())

	p2 := Human{
		name: "charlie",
	}

	fmt.Println(p2)

	p2.SetName("delta")
	fmt.Println(p2.Name())

	// Output:
	// &{bob}
	// alice
	// {charlie}
	// delta
}

type Person struct {
	Firstname string
	Surname   string
}

// String is an implementation of Stringer interface method
func (p Person) String() string {
	return fmt.Sprintf("Firstname: %s Surname: %s", p.Firstname, p.Surname)
}

func Example_person() {
	p := Person{
		Firstname: "John",
		Surname:   "Doe",
	}

	fmt.Println(p)

	// Output:
	// Firstname: John Surname: Doe
}

type Address struct {
	Street string
}

type PersonAddress struct {
	Person
	Address
}

func Example_personAddress() {
	pa := PersonAddress{
		Person{
			Firstname: "John",
			Surname:   "Doe",
		},
		Address{
			Street: "First street",
		},
	}

	fmt.Println(pa) // The Person String() method will be called

	// Output:
	// Firstname: John Surname: Doe
}

type Location struct {
	LocationID string
}

type PersonLocation struct {
	Person
	Location
}

func (p PersonLocation) String() string {
	return fmt.Sprintf("Person: %s %s is in location: %s", p.Firstname, p.Surname, p.LocationID)
}

func Example_personLocation() {
	pl := PersonLocation{
		Person{
			Firstname: "John",
			Surname:   "Doe",
		},
		Location{
			LocationID: "ID-1234",
		},
	}

	fmt.Println(pl)

	// Output:
	// Person: John Doe is in location: ID-1234
}
