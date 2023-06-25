package main

func DivideByZero(numerator int) float64 {
	return float64(numerator / (numerator - 2))
}
