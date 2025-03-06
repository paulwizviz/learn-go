package reader

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
		fmt.Println(n, err, string(buf[:n]))
		if err == io.EOF {
			break
		}
	}

	// Output:
	// 4 <nil> Hell
	// 4 <nil> o Wo
	// 3 <nil> rld
	// 0 EOF

}
