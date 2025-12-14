package basic

import "testing"

func TestSumDigits(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"Positive number", 1234, 10},
		{"Negative number", -55, 10},
		{"Zero", 0, 0},
		{"Single digit", 7, 7},
		{"Large number", 9999, 36},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := SumDigits(tc.input)
			if got != tc.expected {
				t.Errorf("SumDigits(%d) = %d; want %d", tc.input, got, tc.expected)
			}
		})
	}
}