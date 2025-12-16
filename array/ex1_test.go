package array

import "testing"

func TestCalculateSum(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Mảng rỗng",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Toàn số chẵn",
			input:    []int{2, 4, 6},
			expected: 12,
		},
		{
			name:     "Toàn số lẻ",
			input:    []int{1, 3, 5},
			expected: -9,
		},
		{
			name:     "Hỗn hợp chẵn lẻ",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: 3,
		},
		{
			name:     "Có số âm",
			input:    []int{-1, -2, 3, 4},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CalculateSum(tc.input)
			if got != tc.expected {
				t.Errorf("CalculateSum(%v) = %d; want %d", tc.input, got, tc.expected)
			}
		})
	}
}
