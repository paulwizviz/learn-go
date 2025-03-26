package maps

import (
	"fmt"
)

func Example_simpleMap() {
	data := map[string]int{"A": 123, "B": 456}
	fmt.Println(data["A"])
	// Output:
	// 123
}

func Example_makeMap() {
	data := make(map[string]int)
	fmt.Println(len(data))
	data["A"] = 1234
	fmt.Println(data["A"])
	// Output:
	// 0
	// 1234
}

func Example_deleteMapItem() {
	data := map[string]int{"A": 1234, "B": 456}
	fmt.Println(len(data))
	delete(data, "A")
	fmt.Println(len(data))
	// Output:
	// 2
	// 1
}

func Example_accessingNonExistenceItem() {
	data1 := map[string]int{"A": 1234, "B": 456}
	fmt.Println(data1["C"])

	data2 := map[int]string{1: "abc", 2: "efg"}
	fmt.Println(data2[3])

	value, ok := data2[1]
	fmt.Println(ok)
	fmt.Println(value)

	// Output:
	// 0
	//
	// true
	// abc

}

func Example_rangingMaps() {
	data := map[string]int{"A": 1234, "B": 456}
	for k, v := range data {
		switch v {
		case 1234, 456:
			fmt.Println(v)
		default:
			fmt.Println("Value not found")
		}

		switch k {
		case "A", "B":
			fmt.Println(k)
		default:
			fmt.Println("Key not found")
		}
	}
	// Output:
	// 1234
	// A
	// 456
	// B
}

type s struct {
	field int
}

func Example_structKeys() {

	k1 := s{1}
	k2 := s{2}

	m := map[s]string{k1: "hello", k2: "world"}
	v1 := m[k1]
	fmt.Println(v1)

	kp1 := &s{1}
	v2 := m[*kp1]
	fmt.Println(v2)

	// Output:
	// 	hello
	// hello
}
