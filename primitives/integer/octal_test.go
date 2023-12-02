package integer

import "fmt"

func Example_octal() {
	octal0 := 00
	octal1 := 01
	octal2 := 02
	octal3 := 03
	octal4 := 04
	octal5 := 05
	octal6 := 06
	octal7 := 07
	octal10 := 010

	fmt.Printf("Octal0 %v\n", octal0)
	fmt.Printf("Octal1 %v\n", octal1)
	fmt.Printf("Octal2 %v\n", octal2)
	fmt.Printf("Octal3 %v\n", octal3)
	fmt.Printf("Octal4 %v\n", octal4)
	fmt.Printf("Octal5 %v\n", octal5)
	fmt.Printf("Octal6 %v\n", octal6)
	fmt.Printf("Octal7 %v\n", octal7)
	fmt.Printf("Octal10 %v\n", octal10)

	fmt.Printf("Octal0 Literal %0o\n", octal0)
	fmt.Printf("Octal1 Literal %0o\n", octal1)
	fmt.Printf("Octal2 Literal %0o\n", octal2)
	fmt.Printf("Octal3 Literal %0o\n", octal3)
	fmt.Printf("Octal4 Literal %0o\n", octal4)
	fmt.Printf("Octal5 Literal %0o\n", octal5)
	fmt.Printf("Octal6 Literal %0o\n", octal6)
	fmt.Printf("Octal7 Literal %0o\n", octal7)
	fmt.Printf("Octal10 Literal %0o", octal10)

	// Output:
	// Octal0 0
	// Octal1 1
	// Octal2 2
	// Octal3 3
	// Octal4 4
	// Octal5 5
	// Octal6 6
	// Octal7 7
	// Octal10 8
	// Octal0 Literal 0
	// Octal1 Literal 1
	// Octal2 Literal 2
	// Octal3 Literal 3
	// Octal4 Literal 4
	// Octal5 Literal 5
	// Octal6 Literal 6
	// Octal7 Literal 7
	// Octal10 Literal 10
}
