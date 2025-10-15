package reminder_sorting

import "sort"

func remainderSorting(strArr []string) []string {
	mod := 3

	sort.Slice(strArr, func(i, j int) bool {
		return len(strArr[i])%mod < len(strArr[j])%mod || (len(strArr[i])%mod == len(strArr[j])%mod && strArr[i] < strArr[j])
	})

	return strArr
}
