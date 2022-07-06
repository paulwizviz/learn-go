package main

import "fmt"

// When you run this routine you will get this
//
// world
// panic: panic called
//

func main() {
	defer fmt.Println("world")
	panic("panic called")
	fmt.Println("hello")
}
