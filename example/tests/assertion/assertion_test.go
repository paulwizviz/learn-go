package assertion

import (
	"reflect"
	"testing"
)

// echoFunc is a function under test
func echoFunc(input int) int {
	return input
}

func TestAssertion(t *testing.T) {
	input := 1
	expected := 1

	actual := echoFunc(input)
	if expected != actual {
		t.Fatalf("Expected: %v Actual: %v", expected, actual)
	}
}

func TestAssertPointers(t *testing.T) {
	expectedValue := 1
	pExpected := &expectedValue
	actualValue := 1
	pActual := &actualValue

	if pExpected == pActual {
		t.Fatalf("Expected: %v Actual: %v", pExpected, pActual)
	}

	if !reflect.DeepEqual(pExpected, pActual) {
		t.Fatalf("Expected: %v Actual: %v", *pExpected, *pActual)
	}
}

type nested struct {
	value int
}

type data struct {
	value   int
	message string
	nested  nested
}

func TestAssertStruct(t *testing.T) {
	value := 1
	message := "testing"

	expected := data{
		value:   value,
		message: message,
		nested: nested{
			value: 1,
		},
	}

	actual := data{
		value:   value,
		message: message,
		nested: nested{
			value: 1,
		},
	}

	if expected != actual {
		t.Fatalf("Expected: %v Actual: %v", expected, actual)
	}

	if !reflect.DeepEqual(&expected, &actual) {
		t.Fatalf("Expected: %v Actual: %v", expected, actual)
	}
}
