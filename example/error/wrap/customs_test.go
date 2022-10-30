package wrap

import (
	"fmt"
)

type myError struct {
	Description string
	Err         error
}

func (e myError) Error() string {
	return e.Description
}

func myErrorFunc() error {
	return &myError{
		Description: "my error",
	}
}

func Example_myError() {
	err := myErrorFunc()
	if e, ok := err.(*myError); ok {
		fmt.Println(e)
	}

	// Output:
	// my error
}

type underlyingError struct {
	Description string
}

func (u underlyingError) Error() string {
	return u.Description
}

var uErr = &underlyingError{
	Description: "underlying error",
}

func underlyingFunc() error {
	return &myError{
		Description: "my error",
		Err:         uErr,
	}
}

func Example_underlyingError() {
	err := underlyingFunc()
	if e, ok := err.(*myError); ok && e.Err == uErr {
		fmt.Println(e.Err)
	}

	// Output:
	// underlying error
}
