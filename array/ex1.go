package array

func CalculateSum(numbers []int) int {
	result := 0

	for _, v := range numbers {
		if v%2 == 0 {
			result += v
		} else {
			result -= v
		}
	}

	return result
}
