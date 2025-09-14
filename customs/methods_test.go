package customs

import "fmt"

// FloatNumber is an alias for float64
type FloatNumber float64

func (n FloatNumber) Increment() {
	n++ // This will not mutate original value
}

// Pointer type receiver to mutate underlying value
func (n *FloatNumber) IncreaseWith(num FloatNumber) {
	*n = *n + num //
}

func Example_floatNumber() {
	var n FloatNumber = 1.234
	fmt.Println("Original value: ", n)

	n.Increment()
	fmt.Println("After increment: ", n)

	n.IncreaseWith(4.0)
	fmt.Println("Increase by 4.0: ", n)

	var n1 *FloatNumber = &n
	n.IncreaseWith(10.00)
	fmt.Println("Increase by 10.00", *n1)

	// Output:
	// Original value:  1.234
	// After increment:  1.234
	// Increase by 4.0:  5.234
	// Increase by 10.00 15.234
}

type Real struct {
	value float32
}

func (r Real) DoSomething() {
}

func (r *Real) DoAnother() {
	fmt.Printf("Hello: %v\n", r.value)
}

func Example_methodExpression() {

	r := Real{
		value: 1.24,
	}

	fmt.Printf("%T\n", r.DoSomething)
	fmt.Printf("%T\n", r.DoAnother)
	fmt.Printf("%T\n", (Real).DoSomething)
	fmt.Printf("%T\n", (*Real).DoAnother)

	f := (*Real).DoAnother
	f1 := r.DoAnother

	f(&r)
	f1()

	// Output:
	// func()
	// func()
	// func(structs.Real)
	// func(*structs.Real)
	// Hello: 1.24
	// Hello: 1.24
}

// Integer is an alias of int
type Integer int

func (i *Integer) Increment() {
	fmt.Println("increment by 1")
}

func (i *Integer) IncrementWith(integer int) {
	if i == nil {
		panic("(*Integer).Increment is nil")
	}
	*i = *i + Integer(integer) // This will panic if i is nil
}

func Example_numberPtr() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	var i *Integer = nil
	i.Increment()       // This does not panic
	i.IncrementWith(10) // This will panic and fail

	// Output:
	// increment by 1
	// Recovered from panic: (*Integer).Increment is nil
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
