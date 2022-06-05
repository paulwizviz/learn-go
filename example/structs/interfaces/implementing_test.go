package interfaces_test

import "fmt"

// Interfaces for movable object
type movableObject interface {
	move() string
	stop() string
}

// Human is a movable object
type human struct{}

func (h *human) move() string {
	return "Human move"
}

func (h *human) stop() string {
	return "Human stop"
}

// Car is a movable object
type car struct{}

func (c car) move() string {
	return "Car move"
}

func (c car) stop() string {
	return "Car stop"
}

// Animal
type animal struct {
	movableObject
}

func (a animal) move() string {
	return "Animal move"
}

// Services
func move(p movableObject) string {
	switch p.(type) {
	case *human, car:
		return fmt.Sprintf("%v %v", p.stop(), p.move())
	case *animal:
		// p.stop() will panic
		return p.move()
	}
	return ""
}

func Example_movableHuman() {
	h := &human{}
	result := move(h)
	fmt.Println(result)
	// Output:
	// Human stop Human move
}

func Example_movableCar() {
	c := car{}
	result := move(c)
	fmt.Println(result)
	// Output:
	// Car stop Car move
}

func Example_movableAnimal() {
	a := animal{}
	result := move(a)
	fmt.Println(result)
	// Output:
	// Animal move
}
