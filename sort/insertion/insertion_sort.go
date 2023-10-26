package insertion

func insertionSort(arr []int) {
	for right := 1; right < len(arr); right++ {
		pointer := right
		for left := right - 1; left >= 0; left-- {
			if arr[pointer] < arr[left] {
				arr[pointer], arr[left] = arr[left], arr[pointer]
			}
			pointer--
		}
	}
}

func insertionSortOptimized(arr []int) {
	for right := 1; right < len(arr); right++ {
		current := arr[right]
		left := right - 1

		// Use a while loop to find the correct insertion point
		for left >= 0 && arr[left] > current {
			arr[left+1] = arr[left]
			left--
		}

		// Insert the current element into the correct position
		arr[left+1] = current
	}
}
