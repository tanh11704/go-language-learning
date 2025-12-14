package basic

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"Zero input", 0, 1},
		{"Small number", 1, 1},
		{"Standard case", 5, 120},
		{"Larger number", 10, 3628800},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Factorial(tc.input); got != tc.expected {
				t.Errorf("Factorial(%d) = %d; want %d", tc.input, got, tc.expected)
			}

			if got := FactorialRecursive(tc.input); got != tc.expected {
				t.Errorf("FactorialRecursive(%d) = %d; want %d", tc.input, got, tc.expected)
			}
		})
	}
}
