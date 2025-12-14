package basic

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	if a < 0 {
		return -a
	}
	return a
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}
