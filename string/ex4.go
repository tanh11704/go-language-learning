package string

import (
	"strconv"
)

func RunLengthEncoding(s string) string {
	result := ""
	left, right := 0, 0
	char := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != char {
			result += string(char)
			if left < right {
				result += strconv.Itoa(right - left + 1)
			}
			left = i
			char = s[i]
		} else {
			right = i
		}
	}

	result += string(char)
	if len(s)-left > 1 {
		result += strconv.Itoa(len(s) - left)
	}
	return result
}
