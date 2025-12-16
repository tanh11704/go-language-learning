package string

import "strings"

func RepeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 {
			if strings.Repeat(s[0:i], len(s)/(i)) == s {
				return true
			}
		}
	}

	return false
}
