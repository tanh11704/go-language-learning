package string

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"aabaa", true},
		{"z", true},
		{"", true},
		{"Anna", false},
	}

	for _, tc := range testCases {
		t.Run("test", func(t *testing.T) {
			got := IsPalindrome(tc.input)

			if got != tc.expected {
				t.Errorf("IsPalindrome(%s) = %t, want %t", tc.input, got, tc.expected)
			}
		})
	}
}
