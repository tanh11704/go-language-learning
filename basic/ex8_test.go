package basic

import "testing"

func TestIntegerToRoman(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{"Smallest number", 1, "I"},

		{"Subtraction case 4", 4, "IV"},
		{"Subtraction case 9", 9, "IX"},
		{"Complex subtraction", 58, "LVIII"},
		{"Year 1994", 1994, "MCMXCIV"},
		{"Largest number", 3999, "MMMCMXCIX"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IntegerToRoman(tc.input)

			if got != tc.expected {
				t.Errorf("IntegerToRoman(%d) = %s, want %s", tc.input, got, tc.expected)
			}
		})
	}
}
