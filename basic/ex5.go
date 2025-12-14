package basic

func FactorialRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

func Factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
