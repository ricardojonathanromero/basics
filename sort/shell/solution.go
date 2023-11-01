package shell

func shellSort(arr []int) {
	gap := calculateGap(arr)

	// while gap > 0
	for gap > 0 {
		// while selector < length
		for selector := gap; selector < len(arr); selector++ {
			value := arr[selector] // save value in case the value of selector index change
			pointer := selector    // save selector index
			// if pointer greater than gap - 1 and value of pointer minus gap is greater than or equals to value saved
			for pointer > gap-1 && arr[pointer-gap] >= value { // pointer index is equals to pointer minus gap
				arr[pointer] = arr[pointer-gap] // pointer index is equals to the pointer minus gap index
				pointer -= gap                  // pointer = pointer - gap
			}

			arr[pointer] = value // assign value of selector to pointer
		}
		gap = (gap - 1) / 3
	}
}

func calculateGap(arr []int) int {
	var gap int
	for gap < len(arr)/3 {
		gap = gap*3 + 1
	}
	return gap
}
