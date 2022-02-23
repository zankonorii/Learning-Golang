package number_to_string

func NumberToString(x int) string {
	temp := ""
	//for ; x >= 0; x /= 10 {
	l := x % 10
	switch l {
	case 1:
		temp += "one"
		break
	case 2:
		temp += "two"
		break
	case 3:
		temp += "three"
		break
	case 4:
		temp += "four"
		break
	case 5:
		temp += "five"
		break
	case 6:
		temp += "six"
		break
	case 7:
		temp += "seven"
		break
	case 8:
		temp += "eight"
		break
	case 9:
		temp += "nine"
		break

	}
	temp += "  "
	//}
	return temp
}
