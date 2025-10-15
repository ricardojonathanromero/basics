package divisible_sum_pairs

func divisibleSumPairs(n int32, k int32, ar []int32) int {
	var count int

	remainder := make([]int, k)

	for _, num := range ar {
		mod := num % k
		idx := (k - mod) % k
		count += remainder[idx]
		remainder[mod]++
	}

	return count
}
