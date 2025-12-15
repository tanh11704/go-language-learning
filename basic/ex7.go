package basic

func ReverseBinary(n int) int {
	m := 0

	for n > 0 {
		bit := n % 2
		m = m*2 + bit
		n /= 2
	}

	return m
}
