package person

type Student struct {
	Person
	GPA float64
}

func (s Student) Classify() string {
	if s.GPA >= 3.6 {
		return "Excellent"
	}
	if s.GPA >= 3.2 {
		return "Good"
	}
	if s.GPA >= 2.0 {
		return "Average"
	}
	return "Poor"
}
