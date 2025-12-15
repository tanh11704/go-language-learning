package basic

func GetMonthName(month int) string {
	switch month {
	case 1:
		return "Tháng 1"
	case 2:
		return "Tháng 2"
	case 3:
		return "Tháng 3"
	case 4:
		return "Tháng 4"
	case 5:
		return "Tháng 5"
	case 6:
		return "Tháng 6"
	case 7:
		return "Tháng 7"
	case 8:
		return "Tháng 8"
	case 9:
		return "Tháng 9"
	case 10:
		return "Tháng 10"
	case 11:
		return "Tháng 11"
	case 12:
		return "Tháng 12"
	default:
		return "Lỗi"
	}
}
