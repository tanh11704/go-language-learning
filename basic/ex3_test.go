package basic

import (
	"reflect"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected []int
	}{
		{"Standard case", 600, []int{2, 2, 2, 3, 5, 5}},
		{"Prime number", 13, []int{13}},
		{"Square number", 9, []int{3, 3}},
		{"Smallest prime", 2, []int{2}},
		{"Number 1", 1, []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := PrimeFactors(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("PrimeFactors(%d) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}
