package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	calc := Calculator{}
	result := calc.Add(2, 2)
	if result != 4 {
		t.Errorf("Expected 4, got %f", result)
	}
}

func TestSubtract(t *testing.T) {
	calc := Calculator{}
	result := calc.Subtract(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, got %f", result)
	}
}

func TestMultiply(t *testing.T) {
	calc := Calculator{}
	result := calc.Multiply(10, 5)
	if result != 50 {
		t.Errorf("Expected 50, got %f", result)
	}
}

func TestDivide(t *testing.T) {
	calc := Calculator{}
	result, err := calc.Divide(10, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Expected 5, got %f", result)
	}

	_, err = calc.Divide(10, 0)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}

func TestPercentage(t *testing.T) {
	calc := Calculator{}
	result := calc.Percentage(200, 10)
	if result != 20 {
		t.Errorf("Expected 20, got %f", result)
	}
}
