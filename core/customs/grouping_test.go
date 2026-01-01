package customs

import "fmt"

// Cat is an abstraction of feline types
type Cat struct{}

func (c Cat) IsAnimal() {
	fmt.Println("Cat is an animal")
}

func (c Cat) SoundOut() {
	fmt.Println("meow")
}

func (c Cat) Walk() {
	fmt.Printf("On four paws")
}

// Otter can sound out and swim
type Otter struct{}

func (o Otter) IsAnimal() {
	fmt.Println("Otter is an animal")
}

func (o Otter) SoundOut() {
	fmt.Println("squeak")
}

func (o Otter) Walk() {
	fmt.Println("Otter is walking")
}

func (o Otter) Swim() {
	fmt.Println("Otter is swimming")
}

// Fish cannot sound out but can swim
type Fish struct {
}

func (f Fish) IsAnimal() {
	fmt.Println("Fish is an animal")
}

func (f Fish) Swim() {
	fmt.Println("Fish is swimming")
}

// Animal is an interface of concrete
// types with IsAnimal method
type Animal interface {
	IsAnimal()
}

func Example_groupAnimals() {
	animals := []Animal{
		Cat{},
		Otter{},
		Fish{},
	}

	for _, a := range animals {
		a.IsAnimal()
	}

	// Output:
	// Cat is an animal
	// Otter is an animal
	// Fish is an animal
}

// SwimingAnimal is a type of animal that
// can swim
type SwimingAnimal interface {
	Animal
	Swim()
}

func Example_swimmingAnimal() {
	swimmers := []SwimingAnimal{
		Otter{},
		Fish{},
	}

	for _, s := range swimmers {
		fmt.Println("=====")
		s.IsAnimal()
		s.Swim()
	}

	// Output:
	// =====
	// Otter is an animal
	// Otter is swimming
	// =====
	// Fish is an animal
	// Fish is swimming
}
