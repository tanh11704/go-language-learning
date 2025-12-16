package person

import (
	"testing"
)

func TestStudent(t *testing.T) {
	s := Student{
		Person: Person{
			Name:  "Tran Van B",
			Age:   22,
			Email: "student@school.com",
		},
		GPA: 3.8,
	}

	t.Run("Check Embedded Method", func(t *testing.T) {
		if !s.IsValidEmail() {
			t.Error("Student should inherit methods from Person")
		}

		if s.Name != "Tran Van B" {
			t.Error("Should access Name directly")
		}
	})

	t.Run("Check Classification", func(t *testing.T) {
		testCases := []struct {
			gpa      float64
			expected string
		}{
			{4.0, "Excellent"},
			{3.6, "Excellent"},
			{3.5, "Good"},
			{3.2, "Good"},
			{2.5, "Average"},
			{2.0, "Average"},
			{1.9, "Poor"},
			{0.0, "Poor"},
		}

		for _, tc := range testCases {
			tempStudent := Student{GPA: tc.gpa}
			got := tempStudent.Classify()

			if got != tc.expected {
				t.Errorf("GPA %.1f: got %s; want %s", tc.gpa, got, tc.expected)
			}
		}
	})
}
