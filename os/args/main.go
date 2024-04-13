package main

import (
	"flag"
	"log"
	"os"
)

func main() {

	nameFlag := flag.String("name", "value", "An example flag")
	log.Printf("Type: %T Value: %v", nameFlag, *nameFlag)

	args := os.Args
	for index, arg := range args {
		log.Printf("Index: %v Type: %T Value: %v", index, arg, arg)
	}
}
