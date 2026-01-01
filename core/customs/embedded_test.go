package customs

import (
	"fmt"
)

type Thing interface {
	DoSomething() string
	SoundOut() string
}

type LivingThing struct {
	Thing
}

func Example_livingThing() {
	defer func() {
		err := recover()
		if err != nil { //catch
			fmt.Printf("Runtime error: %v", err)
		}
	}()
	var t Thing = LivingThing{}
	t.DoSomething()
	t.SoundOut()
	// Output:
	// Runtime error: runtime error: invalid memory address or nil pointer dereference
}

type HumanBeing struct {
	Thing
}

func (HumanBeing) DoSomething() string {
	return "Walking"
}
func (HumanBeing) SoundOut() string {
	return "Say Hello"
}

func Example_humanBeing() {
	p := HumanBeing{}
	fmt.Println(p.DoSomething())
	fmt.Println(p.SoundOut())

	// Output:
	// Walking
	// Say Hello
}

type SpanishIndividual struct {
	Thing
}

func (s SpanishIndividual) SoundOut() string {
	return "Ola"
}

func Example_spanishIndividual() {
	s := SpanishIndividual{}
	p := HumanBeing{
		s,
	}
	fmt.Println(p.SoundOut())
	fmt.Println(p.Thing.SoundOut())

	// Output:
	// Say Hello
	// Ola
}
