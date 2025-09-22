package writer

import (
	"bytes"
	"fmt"
	"io"
)

func Example_fprint() {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Size: %d MB.", 85)
	s := buf.String()
	fmt.Println(s)

	// Output:
	// Size: 85 MB.
}

func Example_writetobuf() {

	buf := bytes.NewBuffer([]byte("efg"))
	n, err := buf.Write([]byte("abc"))
	fmt.Println(n, err, buf.String())
	buf.Next(1)
	fmt.Println(n, err, buf.String())

	// Output:
	// 3 <nil> efgabc
	// 3 <nil> fgabc
}

type CustomWriter struct {
	cache []byte
}

func (c *CustomWriter) Write(p []byte) (n int, err error) {
	c.cache = append(c.cache, p...)
	//n = len(c.cache)
	err = nil
	if string(c.cache[len(c.cache)-1]) == "\n" {
		fmt.Println("--drain cache--")
		c.cache = []byte{}
		return 0, io.EOF
	}
	c.cache = append(c.cache, p...)
	n = len(c.cache)
	return
}

func NewCustomWriter(p []byte) *CustomWriter {
	return &CustomWriter{
		cache: p,
	}
}

func Example_customWriter() {
	c := NewCustomWriter([]byte("abc"))
	n, err := c.Write([]byte("efg"))
	fmt.Println(n, err)
	n, err = c.Write([]byte("hij"))
	fmt.Println(n, err)
	input := "efg\n"
	n, err = c.Write([]byte(input))
	fmt.Println(n, err)

	// Output:
	// 9 <nil>
	// 15 <nil>
	//--drain cache--
	// 0 EOF
}
