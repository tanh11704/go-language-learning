package basic

func Fibonacci(n int) []int {
	factors := make([]int, 0, 10)
	a, b := 1, 1

	for a <= n {
		factors = append(factors, a)
		a, b = b, a+b
	}

	return factors
}
