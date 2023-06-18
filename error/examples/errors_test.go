package examples

import (
	"errors"
	"fmt"
)

var (
	errInvalidDenominator = errors.New("invalid denominator")
	errLessThenNumberator = errors.New("less than 10")
)

type invalidError struct {
	Description string
}

func (i invalidError) Error() string {
	return i.Description
}

func validate(numerator, denominator int) error {

	if numerator < 10 && denominator == 0 {
		return &invalidError{
			Description: "invalid error",
		}
	}

	if denominator == 0 {
		return errInvalidDenominator
	}

	if numerator < 10 {
		return errLessThenNumberator
	}

	return nil
}

func Example_errorIs() {
	err := validate(11, 0)
	if errors.Is(err, errInvalidDenominator) {
		fmt.Println(err)
	}

	err = validate(1, 1)
	if errors.Is(err, errLessThenNumberator) {
		fmt.Println(err)
	}

	// Output:
	// invalid denominator
	// less than 10

}

func Example_errorAs() {

	err := validate(9, 0)
	var invalidErr *invalidError
	if errors.As(err, &invalidErr) {
		fmt.Println(err)
	}

	// Output:
	// invalid error
}

func partition(x, y int) (int, error) {
	err := validate(x, y)
	if err != nil {
		return 0, fmt.Errorf("div error - %w", err)
	}

	return x / y, nil
}

func Example_wrap() {
	_, err := partition(9, 0)
	if errors.Is(err, errInvalidDenominator) {
		fmt.Println(err)
	}

	if errors.Is(err, errLessThenNumberator) {
		fmt.Println(err)
	}

	var e *invalidError
	if errors.As(err, &e) {
		fmt.Println("Wrapped error message: ", err)
		fmt.Println("Unwrapped error message: ", e.Error())
	}

	// Output:
	// Wrapped error message:  div error - invalid error
	// Unwrapped error message:  invalid error
}

func Example_select() {
	_, err := partition(9, 0)
	switch err.(type) {
	case *invalidError:
		fmt.Println(err)
	case error:
		fmt.Println(err)
	default:
		fmt.Println(err)
	}

	// Output:
	// div error - invalid error
}
