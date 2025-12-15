package basic

import "testing"

func TestReverseBinary(t *testing.T) {
	testCases := []struct {
		intput   int
		expected int
	}{
		{23, 29},
		{12, 3},
		{112, 7},
	}

	for _, tc := range testCases {
		t.Run("test", func(t *testing.T) {
			got := ReverseBinary(tc.intput)

			if tc.expected != got {
				t.Errorf("ReverseBinary(%d) = %d, want %d", tc.intput, got, tc.expected)
			}
		})
	}
}
