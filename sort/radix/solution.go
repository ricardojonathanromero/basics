package radix

const n = 10

func radixSort(arr []int) {
	maximum := arr[0]

	for _, number := range arr[1:] {
		if maximum < number {
			maximum = number
		}
	}

	for exp := 1; maximum/exp > 0; exp *= 10 {
		countSort(arr, exp)
	}
}

func countSort(arr []int, exp int) {
	tmp := make([]int, n)
	out := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		tmp[(arr[i]/exp)%10]++
	}

	for i := 1; i < n; i++ {
		tmp[i] += tmp[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		out[tmp[(arr[i]/exp)%10]-1] = arr[i]
		tmp[(arr[i]/exp)%10]--
	}

	for i := range arr {
		arr[i] = out[i]
	}
}

func radixSortOptimized(arr []int) {
	maximum := arr[0]

	for _, number := range arr[1:] {
		if maximum < number {
			maximum = number
		}
	}

	count := make([]int, n)
	out := make([]int, len(arr))

	for exp := 1; maximum/exp > 0; exp *= 10 {
		for i := range count {
			count[i] = 0
		}

		// Count occurrences of each digit
		for _, num := range arr {
			count[(num/exp)%10]++
		}

		// Calculate cumulative counts
		for i := 1; i < n; i++ {
			count[i] += count[i-1]
		}

		// Build the output array
		for i := len(arr) - 1; i >= 0; i-- {
			out[count[(arr[i]/exp)%10]-1] = arr[i]
			count[(arr[i]/exp)%10]--
		}

		// Copy the output back to the original array
		copy(arr, out)
	}
}
