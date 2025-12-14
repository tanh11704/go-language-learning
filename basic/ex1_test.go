package basic

import "testing"

func TestGCD(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Prime numbers", 5, 13, 1},
		{"Prime numbers 2", 13, 23, 1},
		{"Zero case", 0, 2, 2},
		{"Equal numbers", 10, 10, 10},
		{"Negative number", -5, 1, 1},
		{"Both negative numbers", -7, -3, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GCD(tc.a, tc.b)

			if got != tc.expected {
				t.Errorf("GCD(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}

func TestLCM(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Common numbers", 4, 6, 12},
		{"Prime numbers", 3, 5, 15},
		{"Case with 1", 1, 7, 7},
		{"Large numbers", 100, 25, 100},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := LCM(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("LCM(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}
