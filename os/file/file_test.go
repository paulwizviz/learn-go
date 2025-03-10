package file

import (
	"fmt"
	"os"
)

func Example_fileInfo() {
	f, _ := os.Open("./testdata/test.txt")
	info, _ := f.Stat()
	fmt.Printf("Size: %v\n", info.Size())
	fmt.Printf("Mode: %v\n", info.Mode())
	fmt.Printf("Sys: %T\n", info.Sys())

	// Output:
	// Size: 61
	// Mode: -rw-r--r--
	// Sys: *syscall.Stat_t
}

func Example_read() {
	f, _ := os.Open("./testdata/test.txt")
	data := make([]byte, 2)
	n, _ := f.Read(data)
	fmt.Printf("Count: %d Data: %s\n", n, data)
	n, _ = f.Read(data)
	fmt.Printf("Count: %d Data: %s\n", n, data)
	// Output:
	// Count: 2 Data: Th
	// Count: 2 Data: is
}

func Example_readAt() {
	f, _ := os.Open("./testdata/test.txt")
	data := make([]byte, 2)
	n, _ := f.ReadAt(data, 3)
	fmt.Printf("Count: %d Data: %s\n", n, data)
	n, _ = f.ReadAt(data, 3)
	fmt.Printf("Count: %d Data: %s\n", n, data)

	// Output:
	// Count: 2 Data: s
	// Count: 2 Data: s
}
