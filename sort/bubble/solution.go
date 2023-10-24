package bubble

// []int{64, 34, 25, 12, 22, 11, 90},
func bubbleSort(arr []int) []int {
	// Implement the Bubble Sort algorithm to sort the slice.

	// skipp the first element to compare after
	for i := 1; i < len(arr); i++ {
		// iterate form 0 to length - i. this will ignore rest of items once 'i' increments
		for j := 0; j < len(arr)-i; j++ {
			// if next element if less than current element, swipe values
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}

	return arr
}
