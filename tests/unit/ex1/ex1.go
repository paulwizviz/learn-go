package ex1

import "fmt"

func fizzBuzz(input int) string {

	if input == 0 {
		return "0"
	}

	if input%3 == 0 && input%5 == 0 {
		return "FizzBuzz"
	}

	if input%3 == 0 {
		return "Fizz"
	}
	if input%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%v", input)
}

func fizzBuzzBooleans(input int) string {

	isFizz := input%3 == 0
	isBuzz := input%5 == 0

	if input == 0 {
		return "0"
	} else if isFizz && isBuzz {
		return "FizzBuzz"
	} else if isFizz {
		return "Fizz"
	} else if isBuzz {
		return "Buzz"
	} else {
		return fmt.Sprintf("%v", input)
	}
}
