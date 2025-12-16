package array

func FindMissingNumber(arr []int) int {
	n := len(arr) + 1
	expectedSum := n * (n + 1) / 2

	actualSum := 0
	for _, v := range arr {
		actualSum += v
	}

	return expectedSum - actualSum
}
