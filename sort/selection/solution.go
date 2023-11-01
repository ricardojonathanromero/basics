package selection

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minimum := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minimum] {
				minimum = j
			}
		}

		arr[i], arr[minimum] = arr[minimum], arr[i]
	}
}
