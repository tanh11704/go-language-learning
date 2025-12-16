package array

import (
	"reflect"
	"testing"
)

func TestMergeSortedSlices(t *testing.T) {
	testCases := []struct {
		name     string
		arr1     []int
		arr2     []int
		expected []int
	}{
		{
			name:     "Trộn xen kẽ đều",
			arr1:     []int{1, 3, 5},
			arr2:     []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Arr1 toàn số nhỏ hơn Arr2",
			arr1:     []int{1, 2},
			arr2:     []int{3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Arr2 toàn số nhỏ hơn Arr1",
			arr1:     []int{10, 20},
			arr2:     []int{1, 2},
			expected: []int{1, 2, 10, 20},
		},
		{
			name:     "Có slice rỗng",
			arr1:     []int{1, 5, 9},
			arr2:     []int{},
			expected: []int{1, 5, 9},
		},
		{
			name:     "Cả 2 rỗng",
			arr1:     []int{},
			arr2:     []int{},
			expected: []int{},
		},
		{
			name:     "Có phần tử trùng nhau",
			arr1:     []int{1, 3, 3},
			arr2:     []int{2, 3, 4},
			expected: []int{1, 2, 3, 3, 3, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := MergeSortedSlices(tc.arr1, tc.arr2)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("MergeSortedSlices() = %v; want %v", got, tc.expected)
			}
		})
	}
}
