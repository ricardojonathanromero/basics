package merge

import (
	"github.com/ricardojonathanromero/basics/utils/sort"
	"testing"
)

func fuzzMergeSort(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	actual := mergeSort(input)

	// Verify that the array is sorted
	if !sort.IntSorted(actual) {
		t.Errorf("Merge Sort failed to sort the array: %v", actual)
	}
}

func fuzzMergeSortOptimal(t *testing.T, title string, input []int) {
	t.Logf("start %s test", title)
	// Sort the array using Bubble Sort
	actual := mergeSortOptimized(input)

	// Verify that the array is sorted
	if !sort.IntSorted(actual) {
		t.Errorf("Merge Sort failed to sort the array: %v", actual)
	}
}

func FuzzAll(f *testing.F) {
	f.Add("merge_sort_fuzz_test")
	f.Fuzz(func(t *testing.T, s string) {
		input := sort.GenerateRandomIntSlice(1000)
		fuzzMergeSortOptimal(t, s, input)
		fuzzMergeSort(t, s, input)
	})
}
