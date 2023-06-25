package main

import "fmt"

// A Single method interface.
// In this example we have an interface representing all things movable
type Tracker interface {
	Move(from, to string)
}

// TrackerFunc is a function type that implements a Tracker interface
type TrackerFunc func(string, string)

// Move is a method of TrackerFunc
func (m TrackerFunc) Move(from, to string) {
	m(from, to)
}

// TrackMovementOf is an operation to track item using tracker
func TrackMovementOf(item, from, to string, m Tracker) {
	fmt.Print(item)
	fmt.Print(" ")
	m.Move(from, to)

}

func main() {
	var tracker TrackerFunc = func(from, to string) {
		fmt.Printf("traveling from: %s to %s", from, to)
	}
	TrackMovementOf("John", "a", "b", tracker)
}
