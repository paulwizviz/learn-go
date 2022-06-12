package bytestream

import (
	"fmt"
	"io"
	"strings"
)

func Example_stringReader() {

	r := strings.NewReader("Hello World")

	buf := make([]byte, 4)
	for {
		n, err := r.Read(buf)
		fmt.Println(n, err, buf[:n])
		if err == io.EOF {
			break
		}
	}

	// Output:
	// 4 <nil> [72 101 108 108]
	// 4 <nil> [111 32 87 111]
	// 3 <nil> [114 108 100]
	// 0 EOF []

}

type CustomReader struct {
	src    []byte
	length int
	index  int
}

func (c *CustomReader) Read(p []byte) (int, error) {

	remainder := (c.length - c.index)
	if remainder > len(p) {
		copy(p, c.src[c.index:c.index+len(p)])
		c.index = c.index + len(p)
		return len(p), nil
	} else if remainder > 0 && remainder < len(p) {
		copy(p, c.src[c.index:c.index+remainder])
		c.index = c.index + remainder
		return remainder, nil
	} else {
		return 0, io.EOF
	}

}

func NewCustomReader(input []byte) *CustomReader {
	c := &CustomReader{}
	c.src = make([]byte, len(input))
	c.length = copy(c.src, input)
	c.index = 0
	return c
}

func Example_customReader() {

	r := NewCustomReader([]byte("abc"))

	buf := make([]byte, 5)
	for index := 1; index < 10; index++ {
		n, err := r.Read(buf)
		fmt.Println(n, err, buf[:n])
		if err == io.EOF {
			break
		}
	}

	// Output:
	// 3 <nil> [97 98 99]
	// 0 EOF []
}
