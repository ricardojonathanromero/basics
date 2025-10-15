package longest_palindromic

func longestPalindrome(s string) string {
	var start, end int
	for i := 0; i < len(s); i++ {
		even := expandAroundCenter(s, i, i)
		odd := expandAroundCenter(s, i, i+1)

		maxLen := maximum(even, odd)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func maximum(x, y int) int {
	if x > y {
		return x
	}
	return y
}
