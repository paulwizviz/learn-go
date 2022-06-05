package interfaces_test

import "fmt"

type something interface {
	moveto(x, y int)
	moveback()
}

// data is an existing implementation
type plane struct {
	xPrev, yPrev int
	xCurr, yCurr int
}

func (p *plane) moveto(x, y int) {
	p.xPrev, p.yPrev = p.xCurr, p.yCurr
	p.xCurr, p.yCurr = x, y
}

func (p *plane) moveback() {
	p.xCurr, p.yCurr = p.xPrev, p.yPrev
}

func Example_movePlane() {
	p := &plane{
		xPrev: 0,
		yPrev: 0,
	}

	cruise(p, 10, 10)
	fmt.Println(p)

	cruise(p, 20, 11)
	fmt.Println(p)

	// Output:
	// &{0 0 10 10}
	// &{10 10 20 11}

}

type planeWrapper struct {
	plane
}

func cruise(s something, x, y int) {
	s.moveto(x, y)
}

func Example_movePlaneWrapper() {

	p := &planeWrapper{
		plane{
			xPrev: 0,
			yPrev: 0,
			xCurr: 0,
			yCurr: 0,
		},
	}
	cruise(p, 100, 20)
	fmt.Println(p)

	// Output:
	// &{{0 0 100 20}}
}

type mockSomething struct {
	something
}

func (m *mockSomething) moveto(x, y int) {
	fmt.Println(x, y)
}

func moveSomething(s something) {
	s.moveto(1, 1)
}

func Example_moveSomething() {
	ms := &mockSomething{}
	moveSomething(ms)
	// Output:
	//
}
