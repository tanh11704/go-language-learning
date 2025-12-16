package array

import (
	"testing"
)

func TestFindMissingNumber(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Thiếu ở giữa (thứ tự lộn xộn)",
			input:    []int{3, 7, 1, 2, 8, 4, 5},
			expected: 6,
		},
		{
			name:     "Thiếu số đầu tiên (1)",
			input:    []int{2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Thiếu số cuối cùng (n)",
			input:    []int{1, 2, 3, 4},
			expected: 5,
		},
		{
			name:     "Mảng rỗng (n=1)",
			input:    []int{},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := FindMissingNumber(tc.input)
			if got != tc.expected {
				t.Errorf("FindMissingNumber(%v) = %d; want %d", tc.input, got, tc.expected)
			}
		})
	}
}
