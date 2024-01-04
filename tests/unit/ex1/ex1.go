package ex1

import "fmt"

func fizzbuzz(input uint) string {

	if input == 0 {
		return "0"
	}

	if input%3 == 0 && input%5 == 0 {
		return "Fizz Buzz"
	}

	if input%3 == 0 {
		return "Fizz"
	}
	if input%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%v", input)
}
