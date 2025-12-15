package basic

import "testing"

func TestGetMonthName(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{"January", 1, "Tháng 1"},
		{"February", 2, "Tháng 2"},
		{"March", 3, "Tháng 3"},
		{"April", 4, "Tháng 4"},
		{"May", 5, "Tháng 5"},
		{"June", 6, "Tháng 6"},
		{"July", 7, "Tháng 7"},
		{"August", 8, "Tháng 8"},
		{"September", 9, "Tháng 9"},
		{"October", 10, "Tháng 10"},
		{"November", 11, "Tháng 11"},
		{"December", 12, "Tháng 12"},
		{"Zero", 0, "Lỗi"},
		{"Negative number", -1, "Lỗi"},
		{"Over limit", 13, "Lỗi"},
		{"Large number", 100, "Lỗi"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetMonthName(tc.input)

			if got != tc.expected {
				t.Errorf("GetMonthName(%d) = %s; want %s", tc.input, got, tc.expected)
			}
		})
	}
}
