package ex1

import "fmt"

// Mammal is a custom type that includes these method signatures
type Mammal interface {
	Name() string
	Age() uint
	SetAge(a uint)
}

// human is a struct that implicitly implements Mammal interface
// as long as the struct implements the method signatures
type human struct {
	name string
	age  uint
}

func (m human) Name() string {
	return m.name
}

func (m human) Age() uint {
	return m.age
}

func (m *human) SetAge(a uint) {
	m.age = a
}

func NewMammalHuman() Mammal {
	return &human{
		name: "John",
		age:  10,
	}
}

// dog implements all Mammal methods and more
type dog struct {
	name string
	age  uint
}

func (d dog) Name() string {
	return d.name
}

func (d dog) Age() uint {
	return d.age
}

func (d *dog) SetAge(a uint) {
	d.age = a
}

// Extra method
func (d dog) Bark() {
	fmt.Println("Hello")
}

func NewMammalDog() Mammal {
	return &dog{
		name: "Skippy",
		age:  1,
	}
}

func Example() {
	human := NewMammalHuman()
	dog := NewMammalDog()

	mammals := []Mammal{}
	mammals = append(mammals, human)
	mammals = append(mammals, dog)

	for _, m := range mammals {
		fmt.Printf("Type: %T, Name: %s\n", m, m.Name())
	}

	// output:
	// Type: *main.human, Name: John
	// Type: *main.dog, Name: Skippy
}
