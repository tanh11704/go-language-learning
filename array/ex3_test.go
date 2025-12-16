package array

import (
	"reflect"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Mảng số lẻ phần tử",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Mảng số chẵn phần tử",
			input:    []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "Mảng 1 phần tử",
			input:    []int{99},
			expected: []int{99},
		},
		{
			name:     "Mảng rỗng",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ReserveSlice(tc.input)

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("ReverseSlice() = %v; want %v", got, tc.expected)
			}
		})
	}
}
