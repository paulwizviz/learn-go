package ioexample

type CustomReader struct {
	src    []byte
	length int
	index  int
}

func (c *CustomReader) Read(p []byte) (int, error) {
	copy(p, c.src[c.index:c.index+len(p)])
	c.index = len(p) + c.index
	return len(p), nil

}

func NewCustomReader(input []byte) *CustomReader {
	c := &CustomReader{}
	c.src = make([]byte, len(input))
	c.length = copy(c.src, input)
	c.index = 0
	return c
}
