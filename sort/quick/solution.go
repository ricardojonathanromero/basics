package quick

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	piv := partition(arr)
	quickSort(arr[:piv])
	quickSort(arr[piv+1:])
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr []int) int {
	var itemFromLeft, itemFromRight int
	piv := len(arr) - 1
	itemFromRight = len(arr) - 2

	for {
		for itemFromLeft <= itemFromRight && arr[itemFromLeft] < arr[piv] {
			itemFromLeft++
		}

		for itemFromLeft <= itemFromRight && arr[itemFromRight] > arr[piv] {
			itemFromRight--
		}

		if itemFromLeft >= itemFromRight {
			break
		}

		swap(arr, itemFromLeft, itemFromRight)
		itemFromLeft++
		itemFromRight--
	}

	swap(arr, itemFromLeft, piv)
	return itemFromLeft
}

// optimized
func quickSortOptimized(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Use the median of three to select a good pivot
	pivotIndex := medianOfThree(arr)
	_ = arr[pivotIndex]

	// Partition the array and get the pivot's final position
	pivotIndex = partitionOptimized(arr, pivotIndex)

	// Recursively sort the left and right sub-arrays
	quickSort(arr[:pivotIndex])
	quickSort(arr[pivotIndex+1:])
}

func partitionOptimized(arr []int, pivotIndex int) int {
	pivot := arr[pivotIndex]
	left := 0
	right := len(arr) - 1

	// Move the pivot to the rightmost position
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Perform the partitioning
	for i := range arr {
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Move the pivot to its final sorted position
	arr[right], arr[left] = arr[left], arr[right]

	return left
}

func medianOfThree(arr []int) int {
	if len(arr) < 3 {
		return 0
	}

	first, middle, last := 0, len(arr)/2, len(arr)-1

	if arr[first] > arr[middle] {
		arr[first], arr[middle] = arr[middle], arr[first]
	}
	if arr[middle] > arr[last] {
		arr[middle], arr[last] = arr[last], arr[middle]
	}
	if arr[first] > arr[middle] {
		arr[first], arr[middle] = arr[middle], arr[first]
	}

	return middle
}
