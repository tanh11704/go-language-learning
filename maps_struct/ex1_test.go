package maps

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "Câu đơn giản",
			input: "hello world hello go",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
				"go":    1,
			},
		},
		{
			name:  "Nhiều khoảng trắng thừa",
			input: "   hello   world    ",
			expected: map[string]int{
				"hello": 1,
				"world": 1,
			},
		},
		{
			name:     "Chuỗi rỗng",
			input:    "",
			expected: map[string]int{},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			t.Run(tC.name, func(t *testing.T) {
				got := WordCount(tC.input)

				if !reflect.DeepEqual(got, tC.expected) {
					t.Errorf("WordCount(%q) = %v; want %v", tC.input, got, tC.expected)
				}
			})
		})
	}
}
