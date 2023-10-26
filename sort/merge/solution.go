package merge

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	middle := len(arr) / 2
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])

	return merge(left, right)
}

func mergeSortOptimized(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Split the array into two halves
	middle := len(arr) / 2
	left := make([]int, middle)
	right := make([]int, len(arr)-middle)
	copy(left, arr[:middle])
	copy(right, arr[middle:])

	// Recursively sort both halves
	left = mergeSort(left)
	right = mergeSort(right)

	// Merge the sorted halves
	return merge(left, right)
}

func merge(left, right []int) []int {
	// create a new array with size left.length + right.length
	res := make([]int, 0, len(left)+len(right))

	i, j := 0, 0
	// iterate elements
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return res
}
