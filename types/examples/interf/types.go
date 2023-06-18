package interf

// An interface representing all types
type Mammal interface {
	Name() string
	Age() uint
	SetAge(a uint)
}

// An implementation of Mammal this is typically implemented as a private type
type human struct {
	name string
	age  uint
}

func (h human) Name() string {
	return h.name
}

func (h human) Age() uint {
	return h.age
}

// Note: Use a pointer type for method receiver when you are mutating a member variable
func (h *human) SetAge(a uint) {
	h.age = a
}

// Using a constructor to instantiate a mammal instance
func NewHuman(n string, a uint) Mammal {
	return &human{
		name: n,
		age:  a,
	}
}

// Implementation of mammal representing cats
// In this example you could only implement some of the mammal interface. The compiler will not complain. However, should you invoke an unimplemented method, you will get a panic.
type cat struct {
	Mammal
	age uint
}

func (c cat) Age() uint {
	return c.age
}

func NewCat(a uint) Mammal {
	return &cat{
		age: a,
	}
}

// A Single method interface.
// In this example we have an interface representing all things movable
type Moveable interface {
	Move(from, to string)
}

// Implementing a handler for all things moveable
type MoveableHandler func(string, string)

func (m MoveableHandler) Move(from string, to string) {
	m(from, to)
}
