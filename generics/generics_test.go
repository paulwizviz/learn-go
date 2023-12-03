package generics

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func addInts[T constraints.Integer](a, b T) T {
	return a + b
}

func Example_int() {
	result := addInts[uint8](1, 2)
	fmt.Println(result)

	result1 := addInts(-1, 2) // same as add[int](-1,2)
	fmt.Println(result1)

	result2 := addInts[int32](-1, 2)
	fmt.Println(result2)

	// Output:
	// 3
	// 1
	// 1
}

type Number interface {
	constraints.Integer | constraints.Float
}

func addNumbers[N Number](a, b N) N {
	return a + b
}

func Example_number() {

	result := addNumbers[int](1, 2)
	fmt.Println(result)

	result1 := addNumbers[float32](1.0, 1.2)
	fmt.Println(result1)

	// Output:
	// 3
	// 2.2

}

type Collection[N Number] struct {
	content []N
}

func (c *Collection[N]) Length() int {
	return len(c.content)
}

func (c *Collection[N]) Get(i int) N {
	return c.content[i]
}

func NewCollection[N Number](items []N) *Collection[N] {
	return &Collection[N]{
		content: items,
	}
}

func Example_struct() {
	ci := NewCollection([]int{1, 2, 3})
	result1 := ci.Length()
	fmt.Println(result1)

	cf32 := NewCollection([]float32{1.0, 2.3, 3.4})
	for index := 0; index < cf32.Length(); index++ {
		fmt.Println(cf32.Get(index))
	}

	// Output:
	// 3
	// 1
	// 2.3
	// 3.4
}
