package string

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Không có ký tự lặp lại",
			input:    "ABC",
			expected: "ABC",
		},
		{
			name:     "Lặp lại đơn giản",
			input:    "AA",
			expected: "A2",
		},
		{
			name:     "Lặp lại ở giữa",
			input:    "ABBC",
			expected: "AB2C",
		},
		{
			name:     "Nhiều nhóm lặp lại",
			input:    "AABBBCCCC",
			expected: "A2B3C4",
		},
		{
			name:     "Lặp lại ở cuối",
			input:    "ABB",
			expected: "AB2",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := RunLengthEncoding(tC.input)
			if got != tC.expected {
				t.Errorf("RunLengthEncoding(%q) = %q; want %q", tC.input, got, tC.expected)
			}
		})
	}
}
