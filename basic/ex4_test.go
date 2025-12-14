package basic

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected []int
	}{
		{"Standard case 5", 5, []int{1, 1, 2, 3, 5}},
		{"Limit is not Fibonacci", 10, []int{1, 1, 2, 3, 5, 8}},
		{"Smallest limit", 1, []int{1, 1}},
		{"Zero input", 0, []int{}},
		{"Negative input", -5, []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Fibonacci(tc.n)

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("Fibonacci(%d) = %v; want %v", tc.n, got, tc.expected)
			}
		})
	}
}
