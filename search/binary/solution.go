package binary

func binarySearch(arr []int, target int) int {
	// initialize answer
	answer := -1
	// define low and high
	low, high := 0, len(arr)-1

	// iterate elements
	for low <= high {
		middle := (low + high) / 2
		middleValue := arr[middle]

		// evaluate if equals to target
		if middleValue == target {
			return middle // return middle pointer
		} else if middleValue < target {
			low = middle + 1 // skip previous values
		} else {
			high = middle - 1 // skip next values
		}
	}

	return answer
}
