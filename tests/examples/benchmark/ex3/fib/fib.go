package fib

func FibRecur(n int) int {
	if n <= 1 {
		return n
	}
	return FibRecur(n-1) + FibRecur(n-2)
}
