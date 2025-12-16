package array

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "Tăng dần - Tìm thấy ở giữa",
			arr:      []int{1, 3, 5, 7, 9, 11},
			target:   7,
			expected: 3,
		},
		{
			name:     "Tăng dần - Tìm thấy ở đầu",
			arr:      []int{1, 3, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "Tăng dần - Tìm thấy ở cuối",
			arr:      []int{1, 3, 5},
			target:   5,
			expected: 2,
		},
		{
			name:     "Tăng dần - Không tìm thấy",
			arr:      []int{1, 3, 5, 7},
			target:   4,
			expected: -1,
		},
		{
			name:     "Giảm dần - Tìm thấy",
			arr:      []int{20, 15, 10, 5, 0},
			target:   15,
			expected: 1,
		},
		{
			name:     "Giảm dần - Không tìm thấy",
			arr:      []int{20, 15, 10},
			target:   100,
			expected: -1,
		},
		{
			name:     "Mảng rỗng",
			arr:      []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "Mảng 1 phần tử - Tìm thấy",
			arr:      []int{10},
			target:   10,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BinarySearch(tc.arr, tc.target)
			if got != tc.expected {
				t.Errorf("BinarySearch(%v, %d) = %d; want %d", tc.arr, tc.target, got, tc.expected)
			}
		})
	}
}
