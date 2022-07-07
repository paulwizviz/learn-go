package tests

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
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v Actual: %v", expected, actual)
	}
}
