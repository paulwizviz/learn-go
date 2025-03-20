package slice

import "fmt"

// ReverseInt reverses the content of a slice
func ReverseInt(input []int) []int {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func ExampleReverseInt() {
	input := []int{1, 2, 3, 4, 5}
	output := ReverseInt(input)
	fmt.Println(output)
	// Output:
	// [5 4 3 2 1]
}
