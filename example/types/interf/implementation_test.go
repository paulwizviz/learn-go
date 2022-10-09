package interf_test

import (
	"fmt"

	"github.com/paulwizviz/learn-go/example/types/interf"
)

func Example_fullImpl() {
	human := interf.NewHuman("John", 1)
	fmt.Println(human.Name())
	fmt.Println(human.Age())

	human.SetAge(10)
	fmt.Println(human.Age())

	// Output:
	// John
	// 1
	// 10
}

func Example_partialImpl() {
	cat := interf.NewCat(1)
	fmt.Println(cat.Age())
	// cat.Name() // cause panic as

	// Output:
	// 1
}

// A service to trace locations of moveable thing
func TraceMovement(item, from, to string, mf interf.MoveableHandler) {
	fmt.Print(item)
	mf.Move(from, to)
}

func Example_handler() {

	var humanMovement interf.MoveableHandler = func(from, to string) {
		fmt.Printf("traveling from: %v to %v", from, to)
	}

	TraceMovement("John ", "a", "b", humanMovement)

	// Output:
	// John traveling from: a to b

}
