package main

import (
	"log"

	"github.com/paulwizviz/learn-go/example/primitives/init/sandbox"
)

var _ int = generateInt()

func generateInt() int {
	log.Println("Initialising anonymous main int var")
	return 1
}

func init() {
	log.Println("Main package init")
}

func main() {
	log.Println("Start of main")
	defer log.Println("End of main")
	sandbox.Op()
}
