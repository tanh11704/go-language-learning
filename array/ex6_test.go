package array

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Ví dụ cơ bản",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "Mảng có nhiều số trùng nhau",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "Tìm số lớn nhất (k=1)",
			nums:     []int{1, 2, 3},
			k:        1,
			expected: 3,
		},
		{
			name:     "Tìm số nhỏ nhất (k=n)",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputCopy := make([]int, len(tc.nums))
			copy(inputCopy, tc.nums)

			got := FindKthLargest(inputCopy, tc.k)
			if got != tc.expected {
				t.Errorf("FindKthLargest(%v, %d) = %d; want %d", tc.nums, tc.k, got, tc.expected)
			}
		})
	}
}
