package string

import "testing"

func TestSumOfNumbersInString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"Example case", "abc 123 def 33 mn 3.221", 380},
		{"No numbers", "No numbers here", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := SumOfNumbersInString(tc.input)
			if got != tc.expected {
				t.Errorf("SumOfNumbersInString(%q) = %d; want %d", tc.input, got, tc.expected)
			}
		})
	}
}
