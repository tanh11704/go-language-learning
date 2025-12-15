package basic

func IntegerToRoman(n int) string {
	romanNumerals := []struct {
		number int
		roman  string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""
	for i := 0; i < len(romanNumerals) && n > 0; i++ {
		for n >= romanNumerals[i].number {
			result += romanNumerals[i].roman
			n -= romanNumerals[i].number
		}
	}

	return result
}
