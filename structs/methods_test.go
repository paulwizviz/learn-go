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

// Person represents a person entity
type Person struct {
	name string
}

func (p Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func Example_person() {
	p1 := &Person{
		name: "bob",
	}

	fmt.Println(p1)

	p1.SetName("alice")
	fmt.Println(p1.Name())

	p2 := Person{
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
