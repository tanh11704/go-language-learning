package string

import "testing"

func TestRepeatedSubstringPattern(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Simple repeat", "abab", true},
		{"Not repeated", "aba", false},
		{"Multiple repeats", "abcabcabcabc", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := RepeatedSubstringPattern(tc.input)
			if got != tc.expected {
				t.Errorf("RepeatedSubstringPattern(%q) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}
