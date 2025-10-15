package string_integer

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	const maxVal = 2147483648
	var strToInt int
	multiply := 0
	for _, c := range s {
		if multiply == 0 && c == '-' {
			multiply = -1
			continue
		} else if multiply == 0 && c == '+' {
			multiply = 1
			continue
		}

		if c == ' ' {
			continue
		}

		if c < 48 || c > 57 {
			break
		}

		strToInt = (strToInt * 10) + int(c-48)
	}

	if strToInt > maxVal {
		strToInt = maxVal
	}

	if multiply == 0 {
		multiply = 1
	}

	return strToInt * multiply
}
