package interf

import (
	"fmt"
)

type Thing interface {
	DoSomething() string
	SoundOut() string
}

type Human struct {
	Thing
}

func Example_human() {
	defer func() {
		err := recover()
		if err != nil { //catch
			fmt.Printf("Runtime error: %v", err)
		}
	}()
	var t Thing = Human{}
	t.DoSomething()
	t.SoundOut()
	// Output:
	// Runtime error: runtime error: invalid memory address or nil pointer dereference
}

type Person struct {
	Thing
}

func (Person) DoSomething() string {
	return "Walking"
}
func (Person) SoundOut() string {
	return "Say Hello"
}

func Example_person() {
	p := Person{}
	fmt.Println(p.DoSomething())
	fmt.Println(p.SoundOut())

	// Output:
	// Walking
	// Say Hello
}

type Spanish struct {
	Thing
}

func (s Spanish) SoundOut() string {
	return "Ola"
}

func Example_spanish() {
	s := Spanish{}
	p := Person{
		s,
	}
	fmt.Println(p.SoundOut())
	fmt.Println(p.Thing.SoundOut())

	// Output:
	// Say Hello
	// Ola
}
