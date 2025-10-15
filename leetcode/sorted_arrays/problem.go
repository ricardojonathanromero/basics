package sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, len(nums1)+len(nums2))
	i, j, k := 0, 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged[k] = nums1[i]
			i++
		} else {
			merged[k] = nums2[j]
			j++
		}

		k++
	}

	for i < len(nums1) {
		merged[k] = nums1[i]
		i++
		k++
	}

	for j < len(nums2) {
		merged[k] = nums2[j]
		j++
		k++
	}

	middle := (len(merged) - 1) / 2

	if len(merged)%2 == 0 {
		return float64(merged[middle]+merged[middle+1]) / 2
	}

	return float64(merged[middle])
}
