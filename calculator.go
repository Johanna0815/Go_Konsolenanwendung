package main

import (
	"errors"
	"fmt"
)

// Calculator Struct
type Calculator struct{}

// Functions for basic arithmetic operations and percentage calculation
func (c Calculator) Add(a, b float64) float64 {
	return a + b
}

func (c Calculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c Calculator) Multiply(a, b float64) float64 {
	return a * b
}

func (c Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func (c Calculator) Percentage(a, b float64) float64 {
	return (a * b) / 100
}

func main() {
	calc := Calculator{}

	fmt.Println("Addition of 5 and 3:", calc.Add(5, 3))
	fmt.Println("10% of 50:", calc.Percentage(50, 10))
}
