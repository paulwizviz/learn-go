package integer

import "fmt"

func Example_octal() {
	octal0 := 00
	octal1 := 0o1
	octal2 := 0o2
	octal3 := 03
	octal4 := 04
	octal5 := 05
	octal6 := 06
	octal7 := 07
	// Octal8 := 08 this will cause overflow
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
}

func Example_printOctal() {

	o := 0144
	fmt.Printf("Octal %0o Decimal %d\n", o, o)
	var i int = 100
	fmt.Printf("Octal %0o Deciaml %d\n", i, i)
	var u8 uint8 = 100
	fmt.Printf("Octal %0o Deciaml %d\n", u8, u8)
	var u16 uint16 = 100
	fmt.Printf("Octal %0o Deciaml %d\n", u16, u16)
	var u32 uint32 = 100
	fmt.Printf("Octal %0o Deciaml %d\n", u32, u32)
	var u64 uint64 = 100
	fmt.Printf("Octal %0o Deciaml %d\n", u64, u64)

	// Output:
	// Octal 144 Decimal 100
	// Octal 144 Deciaml 100
	// Octal 144 Deciaml 100
	// Octal 144 Deciaml 100
	// Octal 144 Deciaml 100
	// Octal 144 Deciaml 100
}
