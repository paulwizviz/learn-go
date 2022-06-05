package variable_test

import "fmt"

func Example_variables() {

	const cnum int = 10
	// cnum = 2 -- not permitted
	fmt.Printf("The value of cnum is %v\n", cnum)

	var num int = 1
	fmt.Printf("The value of num is %v\n", num)
	fmt.Printf("The address of num is %v\n", &num)
	num = 2
	fmt.Printf("The value of num is %v\n", num)
	fmt.Printf("The address of num is %v\n", &num)

	fmt.Println("---")

	num1 := 1
	fmt.Printf("The value of num1 is %v\n", num1)
	fmt.Printf("The address of num1 is %v\n", &num1)
	num1 = 2
	fmt.Printf("The num1 is %v\n", num1)
	fmt.Printf("The address of num1 is %v\n", &num1)

	fmt.Println("---")

	var num2 int = num1
	fmt.Printf("The value of num1 is %v\n", num1)
	fmt.Printf("The address of num1 is %v\n", &num1)
	fmt.Printf("The value of num2 is %v\n", num2)
	fmt.Printf("The address of num2 is %v\n", &num2)

	fmt.Println("---")

	num3 := num1
	fmt.Printf("The value of num1 is %v\n", num1)
	fmt.Printf("The address of num1 is %v\n", &num1)
	fmt.Printf("The value of num3 is %v\n", num3)
	fmt.Printf("The address of num3 is %v\n", &num3)

	// Output:

}

func Example_pointers() {

	num := 1
	var intptr *int = &num
	fmt.Printf("The value of num is %v\n", num)
	fmt.Printf("The address of num is %v\n", &num)
	fmt.Printf("The value of intpr is %v\n", *intptr)
	fmt.Printf("The address of intpr is %v\n", intptr)

	fmt.Println("---")

	num = 2
	fmt.Printf("The value of num is %v\n", num)
	fmt.Printf("The address of num is %v\n", &num)
	fmt.Printf("The value of intpr is %v\n", *intptr)
	fmt.Printf("The address of intpr is %v\n", intptr)

	fmt.Println("---")

	*intptr = 3
	fmt.Printf("The value of num is %v\n", num)
	fmt.Printf("The address of num is %v\n", &num)
	fmt.Printf("The value of intpr is %v\n", *intptr)
	fmt.Printf("The address of intpr is %v\n", intptr)

	// Output:

}
