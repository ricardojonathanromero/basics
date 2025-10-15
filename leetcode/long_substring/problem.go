package long_substring

func lengthOfLongestSubstring(s string) int {
	//longest := 0
	longest, start := 0, 0
	dictionary := make(map[rune]int)

	for i, letter := range s {
		if idx, exists := dictionary[letter]; exists && idx >= start {
			start = idx + 1
		}
		dictionary[letter] = i

		current := i - start + 1
		if current > longest {
			longest = current
		}
	}

	return longest
}
