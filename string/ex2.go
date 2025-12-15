package string

import (
	"strconv"
	"strings"
)

func SumOfNumbersInString(s string) int {
	temp := []byte(s)

	for i := 0; i < len(s); i++ {
		if temp[i] < '0' || temp[i] > '9' {
			temp[i] = ' '
		}
	}

	s = string(temp)
	arr_s := strings.Split(s, " ")
	result := 0

	for _, str := range arr_s {
		num, err := strconv.Atoi(str)
		if err == nil {
			result += num
		}
	}

	return result
}
