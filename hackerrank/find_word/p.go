package find_word

func findWord(letters []string) string {
	var word string

	// convert list of letters to map
	words := make(map[uint8]uint8)
	occurrences := make(map[uint8]int)
	uniques := make(map[uint8]bool)

	for _, letter := range letters {
		words[letter[0]] = letter[2]
		occurrences[letter[2]]++  // add just second, considering first letter as primary
		uniques[letter[0]] = true // consider letter position 0 as unique
		uniques[letter[2]] = true // consider letter position 2 as unique
	}

	// find letter not appears in occurrences
	var next uint8
	for unique := range uniques {
		if _, exists := occurrences[unique]; !exists {
			next = unique
		}
	}

	// At this point, I already have the primary letter, only follow the route until the last letter
	for {
		word += string(next)
		nxt, exists := words[next]
		if !exists {
			break
		}
		next = nxt
	}

	return word
}
