package zigzag_conversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	result := make([]byte, 0)
	increment := (numRows - 1) * 2

	for row := 0; row < numRows; row++ {
		for i := row; i < len(s); i += increment {
			result = append(result, s[i])

			if row > 0 && row < numRows-1 && i+increment-2*row < len(s) {
				result = append(result, s[i+increment-2*row])
			}
		}
	}

	return string(result)
}
