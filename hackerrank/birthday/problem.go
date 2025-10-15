package birthday

func birthday(s []int32, d int32, m int32) int32 {
	var result int32

	for i := 0; i <= len(s)-int(m); i++ {
		var sum int32
		for _, num := range s[i : i+int(m)] {
			sum += num
		}

		if sum == d {
			result++
		}
	}

	return result
}
