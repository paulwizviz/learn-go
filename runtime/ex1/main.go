package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Printf("Using %d CPUs", numCPU)
}
