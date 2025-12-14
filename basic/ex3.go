package basic

func PrimeFactors(n int) []int {
	factors := make([]int, 0, 10)

	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	return factors
}
