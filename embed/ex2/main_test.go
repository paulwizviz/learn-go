package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
)

var (

	//go:embed data
	f embed.FS
)

func Example_walk() {
	filesystem := fs.FS(f)
	fs.WalkDir(filesystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		} else {
			fmt.Println(path)
			return nil
		}
	})

	// output:
	// .
	// data
	// data/fruits
	// data/fruits/apples.txt
	// data/lang.txt
	// data/words.txt
}

func Example_sub() {
	filesystem := fs.FS(f)
	fsub, err := fs.Sub(filesystem, "data")
	if err != nil {
		fmt.Println(err)
	} else {
		fs.WalkDir(fsub, "fruits", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				fmt.Println("Directory: ", path)
			} else {
				fmt.Printf("Content of file: %v\n", path)
				file, err := fsub.Open(path)
				if err != nil {
					return err
				}
				b, err := io.ReadAll(file)
				if err != nil {
					return err
				}
				fmt.Println(string(b))
			}
			return nil
		})
	}

	// output:
	// Directory:  fruits
	// Content of file: fruits/apples.txt
	// Granny Smith
	// Fiji
	// Pink lady
}
