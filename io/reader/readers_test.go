package reader

import (
	"fmt"
	"io"
)

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
	}
	if remainder > 0 && remainder < len(p) {
		copy(p, c.src[c.index:c.index+remainder])
		c.index = c.index + remainder
		return remainder, nil
	}
	return 0, io.EOF
}

func NewCustomReader(input []byte) *CustomReader {
	c := &CustomReader{}
	c.src = make([]byte, len(input))
	c.length = copy(c.src, input)
	c.index = 0
	return c
}

func Example_customReader() {

	r := NewCustomReader([]byte("abcefghijklm"))

	buf := make([]byte, 5) // Maximum buffer of size
loop:
	for {
		n, err := r.Read(buf)
		fmt.Println(n, err, string(buf[:n]))
		if err == io.EOF {
			break loop
		}
	}

	// Output:
	// 5 <nil> abcef
	// 5 <nil> ghijk
	// 2 <nil> lm
	// 0 EOF
}
