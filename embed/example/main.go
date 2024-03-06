package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
)

var (

	//go:embed data/*
	f embed.FS

	//go:embed data
	c embed.FS

	//go:embed data/words.txt
	data []byte
)

func main() {
	// Extract data from byte slice
	fmt.Println(string(data))

	fmt.Println("----")

	// Extract embeded text as file
	content, err := f.ReadFile("data/fruits.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	fmt.Println("----")

	// This return a file system
	filesystem := fs.FS(c)
	// This return a set of files
	files, err := fs.Sub(filesystem, "data")
	if err != nil {
		log.Fatal(err)
	}

	f, err := files.Open("lang.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(d))
}
