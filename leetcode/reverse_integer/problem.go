package reverse_integer

func reverse(x int) int {
	const maxVal = 2147483648
	var result int

	for x != 0 {
		num := x % 10
		x /= 10
		result = (result * 10) + num
	}

	if result > maxVal || result < (maxVal*-1) {
		return 0
	}

	return result
}
