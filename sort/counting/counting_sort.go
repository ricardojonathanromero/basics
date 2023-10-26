package counting

func countingSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// 1. Determine the Range
	minimum, maximum := arr[0], arr[0]
	for _, num := range arr {
		if num < minimum {
			minimum = num
		}
		if num > maximum {
			maximum = num
		}
	}

	// 2. Count Occurrences => new arr with size = max - min + 1
	pointer := maximum - minimum + 1
	countingArr := make([]int, pointer)

	// 3. Count Frequencies
	for _, num := range arr {
		// This step counts how many times each element appears in the input.
		countingArr[num-minimum]++
	}

	// 4. Calculate Cumulative Counts
	for i := 1; i < len(countingArr); i++ {
		countingArr[i] += countingArr[i-1]
	}

	// 5. Build the Sorted Output
	out := make([]int, len(arr))
	// 5.1 Starting from the end of the input array
	// 7. Repeat for All Elements
	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		idx := num - minimum
		out[countingArr[idx]-1] = num
		// 6. Decrement Counts
		countingArr[idx]--
	}

	// 8. Final Output
	for i := 0; i < len(out); i++ {
		arr[i] = out[i]
	}
}

func countingOptimizedSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	maximum := arr[0]
	for _, num := range arr {
		if num > maximum {
			maximum = num
		}
	}

	countingArray := make([]int, maximum+1)

	for _, num := range arr {
		countingArray[num]++
	}

	sortedArray := make([]int, len(arr))
	index := 0

	for num, count := range countingArray {
		for i := 0; i < count; i++ {
			sortedArray[index] = num
			index++
		}
	}

	for i := 0; i < len(sortedArray); i++ {
		arr[i] = sortedArray[i]
	}
}
